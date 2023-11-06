package main

import (
	"fmt"
	"log"
	"send-email/config"
	"send-email/email"
)

func main() {
	configEnv, err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	client, err := email.NewSimpleMail(configEnv.Mail)
	if err != nil {
		log.Fatal(err)
	}

	emailParams := email.EmailParams{
		From:    configEnv.Mail.From,
		To:      "myemail@abc.com",
		Subject: "Test Send Email",
		Body:    "Hi. Greet from Indonesia",
	}

	e := make(chan error)
	done := make(chan bool)

	go client.Send(e, done, emailParams)

	select {
	case err = <-e:
		log.Fatal(err)
	case <-done:
		fmt.Println("done.")
	}
}
