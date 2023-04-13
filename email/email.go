package email

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

// SendMail 发送邮件 163
func SendMail(from, to, title, content, mailServer, mailServerPort, mailPassword string) (err error) {
	e := email.NewEmail()
	e.From = from
	e.To = []string{to}
	e.Subject = title
	e.HTML = []byte(content)
	return e.SendWithTLS(fmt.Sprintf("%s%s", mailServer, mailServerPort),
		smtp.PlainAuth("", from, mailPassword, mailServer),
		&tls.Config{InsecureSkipVerify: true, ServerName: mailServer})
}
