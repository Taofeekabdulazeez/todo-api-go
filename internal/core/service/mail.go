package service

import (
	"todo-api-go/pkg/config"

	"gopkg.in/gomail.v2"
)

type MailService struct{}

func (s *MailService) SendVerificationEmail(email string, token string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "taofeekabdulazeez2020@gmail.com")
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "Verify your Signup")
	msg.SetBody("text/html", `
            <p style="text-align: center;">
             <a href="https://localhost:8080/auth/email/signup/callback?token=`+token+`"
              style="background-color: #4CAF50; color: white; padding: 12px 20px; 
              text-decoration: none; font-size: 16px; border-radius: 5px;">
              Complete Signup
             </a>
            </p>
        `)

	d := gomail.NewDialer(config.MAIL_HOST, config.MAIL_PORT, config.MAIL_USERNAME, config.MAIL_PASSWORD)

	return d.DialAndSend(msg)
}
