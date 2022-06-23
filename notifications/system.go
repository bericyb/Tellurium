package notifications

import (
	"fmt"
	"os/exec"
	"time"
)

func Alert(os, command string, totalTime time.Duration) {
	if os == "macOS (Darwin)" {
		result := fmt.Sprintf("display notification \"Execution finished in: %s\" with title \"Tellurium\" subtitle \"%s\" sound name \"\"", totalTime.String(), command)
		cmd := exec.Command("osascript", "-e", result)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
	if os == "Linux" {
		cmd := exec.Command("notify-send", "Tellurium", "Execution Finished in: xxx")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}