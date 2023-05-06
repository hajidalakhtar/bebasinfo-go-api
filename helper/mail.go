package helper

import (
	"gopkg.in/gomail.v2"
	"log"
)

// TODO: TARO DI ENV
const CONFIG_SMTP_HOST = "mail.hajidkh.com"
const CONFIG_SMTP_PORT = 465
const CONFIG_SENDER_NAME = "EventZeez <hajidkhc@hajidkh.com>"
const CONFIG_AUTH_EMAIL = "_mainaccount@hajidkh.com"
const CONFIG_AUTH_PASSWORD = "ASDqwe!@#"

func SendMail(to string, subject, message string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)
	//mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
