package utils

import (
	"fmt"
	"net/smtp"
	"visitor-management-system/config"
	"visitor-management-system/model"
)

func SendSubscriptionEmail(user *model.Subscriber) error {
	to := []string{user.Email}

	address := config.GetConfig().SmtpHost + ":" + config.GetConfig().SmtpPort
	fmt.Println(config.GetConfig().Username)
	//build msg
	subject := "Welcome to VMS"
	body := fmt.Sprintf("dear %s ,\n", user.Name)
	body += fmt.Sprintf("Your %s subscription is activated", user.Subscription_type)

	message := fmt.Sprintf("From: %s\r\n", config.GetConfig().Email)
	message += fmt.Sprintf("To: %s\r\n", to)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += fmt.Sprintf("\r\n%s\r\n", body)

	//fmt.Println(message)
	auth := smtp.PlainAuth("", config.GetConfig().Username, config.GetConfig().SmtpPassword, config.GetConfig().SmtpHost)
	fmt.Println(auth)
	// send mail
	err := smtp.SendMail(address, auth, config.GetConfig().Email, to, []byte(message))

	fmt.Println("email processed")
	return err

}
