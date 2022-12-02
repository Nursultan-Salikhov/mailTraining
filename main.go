package main

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()

	m.SetHeader("From", "from_example@yandex.ru")
	m.SetHeader("To", "to_example@mail.ru")
	m.SetHeader("Subject", "go test")

	m.SetBody("text/plain", "This is example text \n Just text")

	// mail host "smtp.mail.ru", port 465
	// yandex host "smtp.yandex.ru", port 465
	// google host "smtp.gmail.com", port 587
	d := gomail.NewDialer("smtp.yandex.ru", 465, "from_example@yandex.ru", "external_password")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Sending message successful")
}
