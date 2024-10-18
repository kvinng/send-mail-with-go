package main

import (
	"fmt"

	"send-mail-with-go/cmd"
	email "send-mail-with-go/pkg"
)

func Oldmain() {
	email.InitConfig()

	smtpConfig, mail := email.ReadConfig()

	err := email.SendMail(smtpConfig, mail)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	cmd.Execute()
}
