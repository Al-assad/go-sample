package main

import (
	"bytes"
	"log"
	"net/smtp"
)

// @desc go smtp 发送邮件

func main() {
	//sendMail()
	sendMailQuick()
}

// 发送邮件
func sendMail() {
	// 设置认证信息
	auth := smtp.PlainAuth(
		"",               // identity
		"sender@163.com", // username
		"123456",         // password
		"sender@163.com", //  host
	)
	// 连接远程 SMTP 服务员器
	client, err := smtp.Dial("smtp.163.com:465")
	if err != nil {
		log.Fatal(err)
	}
	// 登陆认证
	err = client.Auth(auth)
	if err != nil {
		log.Fatal(err)
	}
	// 发送邮件信息
	client.Mail("sender@163.com")        // 邮件 from 信息
	client.Rcpt("yu.1905.yu@yandex.com") // 邮件 to 信息

	// 邮件正文
	wc, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString("Hello MyFriend, This Email is sent from go smtp client.")
	_, err = buf.WriteTo(wc)
	if err != nil {
		log.Fatal(err)
	}
}

// 快速发送邮件
func sendMailQuick() {
	// 设置认证信息
	auth := smtp.PlainAuth(
		"",               // identity
		"sener@163.com",  // username
		"123456",         // password
		"sender@163.com", //  host
	)
	// 连接到服务器, 认证, 设置发件人、收件人、发送的内容, 然后发送邮件
	err := smtp.SendMail(
		"smtp.163.com:465",                // smtp 服务器
		auth,                              // 认证信息
		"sender@163.com",                  // from
		[]string{"yu.1905.yu@yandex.com"}, // to
		[]byte("Hello MyFriend, This Email is sent from go smtp client."), // 邮件内容
	)
	if err != nil {
		log.Fatal(err)
	}
}
