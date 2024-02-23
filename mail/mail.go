package mail

import (
	"burmese_jewellery/config"
	"fmt"
	"net/smtp"
)

var (
	senderMail     = config.Get().Mail.SenderMail
	senderPassword = config.Get().Mail.SenderPassword
	smtpHost       = "smtp.gmail.com"
	smtpPort       = "587"
	subject        = "Burmese Jewellery「OTP for Email Register」"
	bodyTemplate   = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Password Reset</title>
</head>
<body>
	<p>Please use the following OTP code to authorize email register at Burmese Jewellery</p>
	<h3>%s</h3>
	<p>Thank you for using our service</p>
	<p>Burmese Jewllery Team</p>
</body>
</html>
	`
)

type Mail interface {
	SendMail(email string, otp string) error
}

type mail struct{}

var _ Mail = (*mail)(nil)

func NewMail() Mail {
	return &mail{}
}

func (m *mail) SendMail(receiverMail string, otp string) error {
	auth := smtp.PlainAuth("", senderMail, senderPassword, smtpHost)

	if err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		senderMail,
		[]string{receiverMail},
		[]byte(m.msg(otp)),
	); err != nil {
		return err
	}

	return nil
}

func (m *mail) msg(passwordResetLink string) string {
	headers := make(map[string]string)
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"utf-8\""

	message := ""
	for key, value := range headers {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += "\r\n" + fmt.Sprintf(bodyTemplate, passwordResetLink)

	return message
}
