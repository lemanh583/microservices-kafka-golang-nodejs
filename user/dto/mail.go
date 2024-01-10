package dto

type MailTransferData struct {
	To       []string `json:"to"`
	Subject  string   `json:"subject"`
	Template string   `json:"template"`
}
