package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env  string
	Mail Mail
}

type Mail struct {
	From           string
	Cc             string
	Host           string
	Port           string
	Username       string
	Password       string
	ConnectTimeout string
	SendTimeout    string
}

func LoadEnv() (config Config, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		err = fmt.Errorf("failed to load .env %w", err)
		return
	}

	config = Config{
		Env: os.Getenv("ENV"),
		Mail: Mail{
			From:           os.Getenv("MAIL_FROM"),
			Cc:             os.Getenv("MAIL_CC"),
			Host:           os.Getenv("MAIL_SMTP_HOST"),
			Port:           os.Getenv("MAIL_SMTP_PORT"),
			Username:       os.Getenv("MAIL_SMTP_USERNAME"),
			Password:       os.Getenv("MAIL_SMTP_PASSWORD"),
			ConnectTimeout: os.Getenv("MAIL_CONNECT_TIMEOUT"),
			SendTimeout:    os.Getenv("MAIL_SEND_TIMEOUT"),
		},
	}

	return
}
