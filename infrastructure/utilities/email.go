package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	gmail "gopkg.in/gomail.v2"
)


func SendEmail(to, subject, body string) error {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env" , err.Error())
	}

	username := os.Getenv("USER_EMAIL")
	if username == ""{
		log.Panic("No email")
	}
	password := os.Getenv("EMAIL_PASS")
	if password == ""{
		log.Panic("No password")
	}

	m := gmail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gmail.NewDialer("smtp.gmail.com", 587, username, password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func BodyVerify(token string) (string, string) {
	subject := "Email Verification"
	body := fmt.Sprintf(
	`
	<h2>Verify Your Email</h2>
	<hr>
	<p>Click the link below to verify your email:</p>
	<a href="http://localhost:8080/users/verify-email/%s">Verify Email</a>
	`,token)

	return subject,body
}

func BodyForgetPassword(token string) (string, string) {
	subject := "Forget Password"
	body := fmt.Sprintf(
	`
	<h2>Reset Your Password</h2>
	<hr>
	<p>Click the link below to reset your password:</p>
	<a href="http://localhost:8080/users/password-update/%s">Reset Password</a>
	`,token)

	return subject,body
}