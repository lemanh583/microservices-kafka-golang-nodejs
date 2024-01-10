package config

import "os"

type Config struct {
	DomainGateway    string `env:"DOMAIN_GATEWAY"`
	DatabaseHost     string `env:"DATABASE_HOST"`
	DatabaseName     string `env:"DATABASE_NAME"`
	DatabasePort     string `env:"DATABASE_PORT"`
	DatabaseUsername string `env:"DATABASE_USERNAME"`
	DatabasePassword string `env:"DATABASE_PASSWORD"`
	ServerPort       string `env:"SERVER_PORT"`
	MailSender       string `env:"MAIL_SENDER"`
	MailPassword     string `env:"MAIL_PASSWORD"`
	SMTPHost         string `env:"SMTP_HOST"`
	SMTPPort         string `env:"SMTP_PORT"`
	TokenSecretKey   string `env:"TOKEN_SECRET_KEY"`
}

var Cfg *Config

func InitConfig() {
	Cfg = &Config{
		DomainGateway:    os.Getenv("DOMAIN_GATEWAY"),
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabasePort:     os.Getenv("DATABASE_PORT"),
		DatabaseUsername: os.Getenv("DATABASE_USERNAME"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		ServerPort:       os.Getenv("SERVER_PORT"),
		MailSender:       os.Getenv("MAIL_SENDER"),
		MailPassword:     os.Getenv("MAIL_PASSWORD"),
		SMTPHost:         os.Getenv("SMTP_HOST"),
		SMTPPort:         os.Getenv("SMTP_PORT"),
		TokenSecretKey:   os.Getenv("TOKEN_SECRET_KEY"),
	}
}
