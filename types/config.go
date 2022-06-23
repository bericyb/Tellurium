package types

type Config struct {
	PhoneNumber            string `json:"phone_number"`
	Email                  string `json:"email"`
	SystemAlert            bool   `json:"system_alert"`
	SystemAlertEmail       bool   `json:"alert_email"`
	SystemAlertPhoneNumber bool   `json:"alert_phone_number"`
	SendReports            bool   `json:"send_reports"`
	Logs                   bool   `json:"logs"`
	Machine                string `json:"machine"`
}