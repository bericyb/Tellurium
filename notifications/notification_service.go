package notifications

import (
	"te/types"
	"time"
)

var (
	SendAlert      = Alert
	SendTextAlert  = SendAWSTextAlert
	SendEmailAlert = SendAWSEmailAlert
)

func Notify(configuration types.Config, command string, totalTime time.Duration) {
	if configuration.SystemAlert {
		SendAlert(configuration.Machine, command, totalTime)
	}

	if configuration.SystemAlertPhoneNumber {
		SendTextAlert(configuration.PhoneNumber, command, totalTime)
	}

	if configuration.SystemAlertEmail {
		SendEmailAlert(configuration.Email, command, totalTime)
	}
}
