/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	email "send-mail-with-go/pkg"

	"github.com/spf13/cobra"
)

// sendmailCmd represents the sendmail command
var sendmailCmd = &cobra.Command{
	Use:   "sendmail",
	Short: "Send email with Go(lang)",
	Long:  `Send email with file config.yml`,
	Run: func(cmd *cobra.Command, args []string) {

		smtpConfig, mail := email.ReadConfig()

		err := email.SendMail(smtpConfig, mail)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sendmailCmd)
}
