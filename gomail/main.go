package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	gomail "gopkg.in/gomail.v2"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	var smtp_mail = os.Getenv("SMTP_MAIL")
	var smtp_key = os.Getenv("SMTP_KEY")

	if smtp_mail == "" || smtp_key == "" {
		fmt.Println("SMTP credentials not found in environment variables")
		return
	}

	sndMail := gomail.NewMessage()

	sndMail.SetHeader("From", smtp_mail)
	sndMail.SetHeader("To", "robert@mailinator.com")
	sndMail.SetHeader("Subject", "Email Test")
	sndMail.SetBody("text/plain", "Just saying Hiiii!")

	a := gomail.NewDialer("smtp.gmail.com", 587, smtp_mail, smtp_key)

	if err := a.DialAndSend(sndMail); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
