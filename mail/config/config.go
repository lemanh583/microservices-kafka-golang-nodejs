package config

import "os"

type Config struct {
	MailSender   string `env:"MAIL_SENDER"`
	MailPassword string `env:"MAIL_PASSWORD"`
	SMTPHost     string `env:"SMTP_HOST"`
	SMTPPort     string `env:"SMTP_PORT"`
}

var Cfg *Config

func InitConfig() {
	Cfg = &Config{
		MailSender:   os.Getenv("MAIL_SENDER"),
		MailPassword: os.Getenv("MAIL_PASSWORD"),
		SMTPHost:     os.Getenv("SMTP_HOST"),
		SMTPPort:     os.Getenv("SMTP_PORT"),
	}
}
