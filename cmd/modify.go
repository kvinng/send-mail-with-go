/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	email "send-mail-with-go/pkg"

	"github.com/spf13/cobra"
)

var smtpServer, smtpUsername, smtpPassword string
var smtpPort int
var mailFromName, mailFrom, mailSubject, mailBody string
var mailTo []string

// modifyCmd represents the modify command
var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify file config.yml with flags",
	Long:  `Modify config.yml with flags in execution time.`,
	Run: func(cmd *cobra.Command, args []string) {
		smtpConfig, mailConfig, err := email.OtherWayToReadAndParse()

		if err != nil {
			log.Fatalf("Error al leer la configuración: %v", err)
		}
		if err := email.SendMail(smtpConfig, mailConfig); err != nil {
			fmt.Printf("errior %s", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(modifyCmd)
	modifyCmd.Flags().StringVar(&smtpServer, "smtp-server", "", "New STMP server")
	modifyCmd.Flags().StringVar(&smtpUsername, "smtp-username", "", "New STMP user")
	modifyCmd.Flags().StringVar(&smtpPassword, "smtp-password", "", "New STMP password")
	modifyCmd.Flags().IntVar(&smtpPort, "smtp-port", 0, "New STMP port")

	modifyCmd.Flags().StringVar(&mailFromName, "mail-fromName", "", "Email From Name")
	modifyCmd.Flags().StringVar(&mailFrom, "mail-from", "", "Email From")
}
