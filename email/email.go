package email

import (
	"fmt"
	"send-email/config"
	"strconv"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

type EmailParams struct {
	From    string
	To      string
	Subject string
	Body    string
}

type mailClient struct {
	smtp *mail.SMTPClient
}

func NewSimpleMail(config config.Mail) (*mailClient, error) {
	port, err := strconv.Atoi(config.Port)
	if err != nil {
		return nil, err
	}

	connectTimeout, err := time.ParseDuration(config.ConnectTimeout)
	if err != nil {
		return nil, err
	}

	sendTimeout, err := time.ParseDuration(config.SendTimeout)
	if err != nil {
		return nil, err
	}

	dialer := mail.NewSMTPClient()
	dialer.Host = config.Host
	dialer.Port = port
	dialer.Username = config.Username
	dialer.Password = config.Password
	dialer.Encryption = mail.EncryptionSTARTTLS
	dialer.ConnectTimeout = connectTimeout
	dialer.SendTimeout = sendTimeout
	dialer.KeepAlive = false

	client, err := dialer.Connect()
	if err != nil {
		return nil, err
	}

	mailClient := mailClient{
		smtp: client,
	}

	return &mailClient, nil
}

func (c *mailClient) Send(e chan error, done chan bool, params EmailParams) {
	defer close(done)
	defer c.smtp.Close()

	mailer := mail.NewMSG()
	mailer.SetFrom(params.From)
	mailer.AddTo(params.To)
	mailer.SetSubject(params.Subject)
	mailer.AddAlternative(mail.TextPlain, params.Body)

	fmt.Print("process sending email... ")

	err := mailer.Send(c.smtp)
	if err != nil {
		e <- fmt.Errorf("error send email: %s", err.Error())
		return
	}

	fmt.Println("sent.")
}
