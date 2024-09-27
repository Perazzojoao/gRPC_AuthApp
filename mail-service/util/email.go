package util

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

type Message struct {
	From    string
	To      string
	Subject string
	Data    any
	DataMap map[string]any
}

var (
	domain   = os.Getenv("MAIL_DOMAIN")
	user     = os.Getenv("MAIL_USER")
	password = os.Getenv("MAIL_PASSWORD")

	//go:embed templates/verifyMail.html.gohtml
	//go:embed templates/resetPasswordMail.html.gohtml
	mailTemplates embed.FS
)

const (
	VerificationCode = "verification_code"
	ResetPassword    = "reset_password"
)

func (msg *Message) SendGomail(about string) error {
	msg.From = user
	msg.DataMap = map[string]any{
		"message": msg.Data,
	}

	var body bytes.Buffer

	tpl, err := selectTemplate(about)
	if err != nil {
		log.Println(err)
		return err
	}

	tpl.ExecuteTemplate(&body, "body", msg.DataMap)

	m := gomail.NewMessage()
	m.SetHeader("From", msg.From)
	m.SetHeader("To", msg.To)
	m.SetHeader("Subject", msg.Subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(fmt.Sprintf("smtp.%s", domain), 587, user, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func selectTemplate(about string) (*template.Template, error) {
	var tpl *template.Template
	var err error
	switch about {
	case VerificationCode:
		tpl, err = template.New(VerificationCode).ParseFS(mailTemplates, "templates/verifyMail.html.gohtml")

	case ResetPassword:
		tpl, err = template.New(ResetPassword).ParseFS(mailTemplates, "templates/resetPasswordMail.html.gohtml")

	default:
		err = fmt.Errorf("invalid about")
	}

	return tpl, err
}
