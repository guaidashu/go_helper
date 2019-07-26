/**
  create by yy on 2019-07-26
*/

package go_helper

import "github.com/go-gomail/gomail"

// 发送简单 邮件 (内容为字符串)
// Sending a simple email, the content can only be string.
func SendEmail(subject string, message string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "")
	m.SetHeader("To", "")
	// 设置主题
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)
	d := gomail.NewDialer("smtp.qq.com", 587, "", "")
	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
