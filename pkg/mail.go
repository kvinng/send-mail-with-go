package email

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server SMTPServer `yaml:"smtp_server"`
	Email  Email      `yaml:"mail"`
}

type Email struct {
	FromName string   `yaml:"from_name"`
	From     string   `yaml:"from"`
	To       []string `yaml:"to"`
	Subject  string   `yaml:"subject"`
	Body     string   `yaml:"body"`
}

type SMTPServer struct {
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Receive STMPServer & Email struct and return error if unable send email
func SendMail(server SMTPServer, mail Email) error {
	auth := smtp.PlainAuth("", server.Username, server.Password, server.Server)
	message := formatMessageBody(mail)

	err := smtp.SendMail(server.Server+":"+strconv.Itoa(server.Port), auth, mail.From, mail.To, []byte(message))
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func formatMessageBody(mail Email) string {
	return fmt.Sprintf(
		"From: %s <%s>\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"Message-ID: %s\r\n"+
			"Date: %s\r\n"+
			"\r\n%s",
		mail.FromName, mail.From, strings.Join(mail.To, ","), mail.Subject, createMessageID(mail.From), time.Now().Format(time.RFC1123Z), mail.Body)
}

func createMessageID(from string) string {
	return fmt.Sprintf("<%d.%d@%s>", time.Now().UnixNano(), time.Now().Unix(), from)
}

func InitConfig() error {
	fileName := "config.yml"

	config := Config{
		Server: SMTPServer{
			Server:   "smtp.server.foo",
			Port:     587,
			Username: "userSMTP",
			Password: "passwordSMTP",
		},
		Email: Email{
			FromName: "MailFrom",
			From:     "sender@mail.foo",
			To:       []string{"receiver@mail.foo", "receiver2@mail.foo"},
			Subject:  "Email from Golang",
			Body:     "This is sending an email with a program made with Go❤️",
		},
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("[Error] Parse configuration to YAML: %v", err)
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		err = os.WriteFile(fileName, data, 0644)
		if err != nil {
			return fmt.Errorf("[Error] Unable to write in YAML file: %v", err)
		}
		fmt.Printf("File '%s' create successfully.\n", fileName)
	} else {
		fmt.Printf("File '%s' exist. Modify that file.\n", fileName)
	}

	return nil
}

func ReadConfig() (SMTPServer, Email) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	return SMTPServer{
			Server:   viper.GetString("smtp_server.server"),
			Port:     viper.GetInt("smtp_server.port"),
			Username: viper.GetString("smtp_server.username"),
			Password: viper.GetString("smtp_server.password"),
		}, Email{
			FromName: viper.GetString("mail.from_name"),
			From:     viper.GetString("mail.from"),
			To:       viper.GetStringSlice("mail.to"),
			Subject:  viper.GetString("mail.subject"),
			Body:     viper.GetString("mail.body"),
		}
}

func OtherWayToReadAndParse() (SMTPServer, Email, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return SMTPServer{}, Email{}, fmt.Errorf("error al leer el archivo de configuración: %v", err)
	}

	var smtpConfig SMTPServer
	var emailConfig Email

	err := viper.UnmarshalKey("smtp_server", &smtpConfig)
	if err != nil {
		return SMTPServer{}, Email{}, fmt.Errorf("error al decodificar smtp_server: %v", err)
	}

	err = viper.UnmarshalKey("mail", &emailConfig)
	if err != nil {
		return SMTPServer{}, Email{}, fmt.Errorf("error al decodificar mail: %v", err)
	}

	return smtpConfig, emailConfig, nil
}

func SaveConfig(smtp SMTPServer, mail Email) error {
	viper.Set("smtp_server", smtp)
	viper.Set("mail", mail)

	config := viper.AllSettings()
	data, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("error al convertir la configuración a YAML: %v", err)
	}

	err = os.WriteFile("config.yml", data, 0644)
	if err != nil {
		return fmt.Errorf("error al escribir el archivo YAML: %v", err)
	}

	return nil
}
