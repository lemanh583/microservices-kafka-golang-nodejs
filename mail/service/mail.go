package service

import (
	"fmt"
	"learn-microservices-mail/config"
	"net/smtp"
)

type mailService struct{}

type MailService interface {
	SendMail(to []string, subject string, template string) error
}

func NewMailService() *mailService {
	return &mailService{}
}

func (s *mailService) SendMail(to []string, subject string, template string) error {
	message := ""
	smtpHost := config.Cfg.SMTPHost
	smtpPort := config.Cfg.SMTPPort
	from := config.Cfg.MailSender
	password := config.Cfg.MailPassword

	headers := make(map[string]string)
	headers["Subject"] = subject
	for k, v := range headers {
		message += k + ": " + v + "\r\n"
	}
	message += "\r\n" + template

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		fmt.Println("SendMail failed", err)
		return err
	}
	return nil
}
