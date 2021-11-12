package utils

import (
	"log"
	"strconv"
	"time"

	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	mail "github.com/xhit/go-simple-mail/v2"
)

// Establish email server connecton and return an email client
func getEmailClient(env Config) *mail.SMTPClient {

	server := mail.NewSMTPClient()
	port, err := strconv.Atoi(env.MailPort)
	if err != nil {
		logger.Error("StrConError: " + err.Error())
	}

	server.Host = env.MailHost
	server.Port = port
	server.Username = env.MailUserName
	server.Password = env.MailPassword
	server.Encryption = mail.EncryptionSSLTLS //mail.EncryptionTLS
	server.KeepAlive = true
	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second
	// Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second
	smtpClient, err := server.Connect()
	if err != nil {
		logger.Error(konstants.Mail_CON_ERR + err.Error())
		log.Fatal(err)
	}

	return smtpClient
}
