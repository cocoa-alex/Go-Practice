package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

const (
	HOST        = "smtp.163.com"
	SERVER_ADDR = "smtp.163.com:25"
	USER        = "XXXXX@163.com"
	PASSWORD    = "XXXXX"
)

type Email struct {
	to      string "to"
	subject string "subject"
	msg     string "msg"
}

func NewEmail(to, subject, msg string) *Email {
	return &Email{to: to, subject: subject, msg: msg}
}

func SendEmail(email *Email) error {
	author := smtp.PlainAuth("", USER, PASSWORD, HOST)

	sendTo := strings.Split(email.to, ";")

	done := make(chan error, 1024)

	go func() {
		defer close(done)

		for _, v := range sendTo {
			str := strings.Replace("From: "+USER+"~To: "+v+"~Subject: "+email.subject+"~~", "~", "\r\n", -1) + email.msg

			err := smtp.SendMail(SERVER_ADDR, author, USER, []string{v}, []byte(str))

			done <- err
			fmt.Println(done)
		}
	}()

	for i := 0; i < len(sendTo); i++ {
		<-done

		//fmt.Println(res)
	}
	return nil
}

func main() {
	em := NewEmail("3146916511@qq.com;25162251@qq.com", "test", "this is a test email by useing go")

	err := SendEmail(em)

	fmt.Println(err)
}
