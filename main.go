package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"te/types"
	"time"

	"te/notifications"

	"github.com/mitchellh/go-homedir"
)

func main() {

	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	banner()
	expanded, _ := homedir.Expand("~/.tellurium")

	configData, _ := os.Open(expanded)
	defer configData.Close()
	decoder := json.NewDecoder(configData)
	configuration := types.Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		createConfig()
	}

	const (
		defaultConfig   = false
		configUsage     = "Display and edit the configuration"
		alertUsage      = "Send alert to local system"
		phoneUsage      = "Send alert to phone number"
		emailUsage      = "Send alert to email address"
		logUsage        = "Display logs to console"
		sendReportUsage = "Send report to email address"
	)

	var config bool
	flag.BoolVar(&config, "c", defaultConfig, configUsage+" (shorthand)")
	flag.BoolVar(&config, "config", defaultConfig, configUsage)
	flag.BoolVar(&configuration.SystemAlert, "a", configuration.SystemAlert, alertUsage+" (shorthand)")
	flag.BoolVar(&configuration.SystemAlert, "alert", configuration.SystemAlert, alertUsage)
	flag.BoolVar(&configuration.SystemAlertPhoneNumber, "p", configuration.SystemAlertPhoneNumber, phoneUsage+" (shorthand)")
	flag.BoolVar(&configuration.SystemAlertPhoneNumber, "phone", configuration.SystemAlertPhoneNumber, phoneUsage)
	flag.BoolVar(&configuration.SystemAlertEmail, "e", configuration.SystemAlertEmail, emailUsage+" (shorthand)")
	flag.BoolVar(&configuration.SystemAlertEmail, "email", configuration.SystemAlertEmail, emailUsage)
	flag.BoolVar(&configuration.Logs, "l", configuration.Logs, logUsage+" (shorthand)")
	flag.BoolVar(&configuration.Logs, "log", configuration.Logs, logUsage)
	flag.BoolVar(&configuration.SendReports, "r", configuration.SendReports, sendReportUsage+" (shorthand)")
	flag.BoolVar(&configuration.SendReports, "report", configuration.SendReports, sendReportUsage)

	flag.Parse()

	if config {
		createConfig()
	}

	if len(flag.Args()) == 0 {
		os.Exit(0)
	}

	t := time.Now()
	cmd := exec.Command(flag.Args()[0], flag.Args()[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(out.String())

	totalTime := time.Now().Sub(t)

	command := strings.Join(flag.Args(), " ")

	notifications.Notify(configuration, command, totalTime)
}

func createConfig() {
	fmt.Println("Building new Tellurium configuration...")
	fmt.Println("")

	machine := ""
	OperatingSystem := runtime.GOOS
	switch OperatingSystem {
	case "windows":
		machine = ("Windows")
	case "darwin":
		machine = ("macOS (Darwin)")
	case "linux":
		machine = ("Linux")
	default:
		machine = (OperatingSystem)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Phone Number: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.Replace(phone, "\n", "", -1)

	fmt.Print("")

	fmt.Print("Send Phone Number Alerts (y/n): ")
	alertPhone, _ := reader.ReadString('\n')
	alertPhone = strings.Replace(alertPhone, "\n", "", -1)

	fmt.Print("")

	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.Replace(email, "\n", "", -1)

	fmt.Print("")

	fmt.Print("Send Email Alerts (y/n): ")
	alertEmail, _ := reader.ReadString('\n')
	alertEmail = strings.Replace(alertEmail, "\n", "", -1)

	fmt.Print("")

	fmt.Print("System Alert (y/n): ")
	system, _ := reader.ReadString('\n')
	system = strings.Replace(system, "\n", "", -1)

	fmt.Print("")

	fmt.Print("Display Logs (y/n): ")
	logs, _ := reader.ReadString('\n')
	logs = strings.Replace(logs, "\n", "", -1)

	fmt.Print("")

	fmt.Print("Send Analytics (y/n): ")
	analytics, _ := reader.ReadString('\n')
	analytics = strings.Replace(analytics, "\n", "", -1)

	configuration := types.Config{
		PhoneNumber:            phone,
		Email:                  email,
		SystemAlert:            system == "y",
		SystemAlertEmail:       alertEmail == "y",
		SystemAlertPhoneNumber: alertPhone == "y",
		SendReports:            analytics == "y",
		Logs:                   logs == "y",
		Machine:                machine,
	}
	file, err := json.MarshalIndent(configuration, "", "    ")
	if err != nil {
		fmt.Println("Could not marshall config file")
		os.Exit(1)
	}
	expanded, _ := homedir.Expand("~/.tellurium")
	err = ioutil.WriteFile(expanded, file, 0644)
	if err != nil {
		fmt.Println("Error saving config file")
	} else {
		fmt.Println("Config file created")
	}
}

func banner() {
	ascii := fmt.Sprintf(`
-------------------------------------------------------------------------
$$$$$$$$\        $$\ $$\                     $$\                         
\__$$  __|       $$ |$$ |                    \__|                        
   $$ | $$$$$$\  $$ |$$ |$$\   $$\  $$$$$$\  $$\ $$\   $$\ $$$$$$\$$$$\  
   $$ |$$  __$$\ $$ |$$ |$$ |  $$ |$$  __$$\ $$ |$$ |  $$ |$$  _$$  _$$\ 
   $$ |$$$$$$$$ |$$ |$$ |$$ |  $$ |$$ |  \__|$$ |$$ |  $$ |$$ / $$ / $$ |
   $$ |$$   ____|$$ |$$ |$$ |  $$ |$$ |      $$ |$$ |  $$ |$$ | $$ | $$ |
   $$ |\$$$$$$$\ $$ |$$ |\$$$$$$  |$$ |      $$ |\$$$$$$  |$$ | $$ | $$ |
   \__| \_______|\__|\__| \______/ \__|      \__| \______/ \__| \__| \__|
		 
"Just tell me when it's done"                           Tellurium v0.1.0
------------------------------------------------------------------------ `, '`')
	fmt.Println(ascii)
}
