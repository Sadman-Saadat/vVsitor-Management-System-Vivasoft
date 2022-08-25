package utils

import (
	"fmt"
	"net/smtp"
	"visitor-management-system/config"
	//"visitor-management-system/model"
)

func SendEmail(email string, password string, subdomain string, link string, company_name string) error {
	to := []string{email}
	var body string
	address := config.GetConfig().SmtpHost + ":" + config.GetConfig().SmtpPort

	subject := "Subject:Welcome to VMS! \r\n"
	from := fmt.Sprintf("From:%s \r\n", config.GetConfig().Email)
	send_to := fmt.Sprintf("To:%s \r\n", email)
	if link != "" {

		color := "color:red;"
		mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body := fmt.Sprintf("<html><body><h4>Hello %s,</br> You registered an account on PiVisitor.</br> We just need to verify your email address before you can access PiVisitor.</br>Verify your email address</h4></br><span style=%s><a href=%s>Click Here!</a></span></br><h3>Company subdomain:%s</h3><h3>Admin email:%s</h3></body></html>", company_name, color, link, subdomain, email)
		msg := []byte(from + send_to + subject + mime + body)
		auth := smtp.PlainAuth("", config.GetConfig().Username, config.GetConfig().SmtpPassword, config.GetConfig().SmtpHost)
		err := smtp.SendMail(address, auth, config.GetConfig().Email, to, msg)
		return err
	} else {
		body = fmt.Sprintf("Your credentials for login are given below: \n")
		body += fmt.Sprintf("Username: %s \n", email)
		body += fmt.Sprintf("Password: %s\n", password)
	}

	if subdomain != "" {
		body += fmt.Sprintf("Company SubDomain: %s\n", subdomain)
	}

	message := fmt.Sprintf("From: %s\r\n", config.GetConfig().Email)
	message += fmt.Sprintf("To: %s\r\n", to)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += fmt.Sprintf("\r\n%s\r\n", body)

	auth := smtp.PlainAuth("", config.GetConfig().Username, config.GetConfig().SmtpPassword, config.GetConfig().SmtpHost)
	err := smtp.SendMail(address, auth, config.GetConfig().Email, to, []byte(message))

	fmt.Println("email processed")
	return err

}
