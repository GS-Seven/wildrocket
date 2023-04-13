package email

import (
	"fmt"
	"os"
	"testing"
)

func TestSendMail(t *testing.T) {
	from := os.Getenv("MailFrom")
	to := os.Getenv("MailTo")
	mailPassword := os.Getenv("Mail163Pass")
	mailServer := os.Getenv("MailServer")
	mailServerPort := os.Getenv("MailServerPort")
	title := "单元测试，邮件发送"
	content := `能发送过去吗:<h1>" + 123456 + "</h1>`
	err := SendMail(from, to, title, content, mailServer, mailServerPort, mailPassword)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("邮件发送成功...")
}
