package util

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendGridMail(name, email, subject, fileName, token string) {

	from := mail.NewEmail("admin", "admin@unindra.com")
	to := mail.NewEmail(name, email)
	subjectMail := subject
	template := ParseHtml(fileName, map[string]string{
		"to":    email,
		"token": token,
	})

	message := mail.NewSingleEmail(from, subjectMail, to, "", template)
	client := sendgrid.NewSendClient(GodotEnv("SG_API_KEY"))
	client.Send(message)
}
