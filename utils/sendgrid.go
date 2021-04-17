package util

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// type Request struct {
// 	Template string
// }

// func NewRequest(Template string) *Request {
// 	return &Request{
// 		Template: Template,
// 	}
// }

func SendGridMail(name, email, subjectMail, token string) {

	from := mail.NewEmail("admin", "admin@unindra.com")
	to := mail.NewEmail(name, email)
	subject := subjectMail
	template := "xxxx"

	message := mail.NewSingleEmail(from, subject, to, "", template)
	client := sendgrid.NewSendClient(GodotEnv("SG_API_KEY"))
	client.Send(message)
}
