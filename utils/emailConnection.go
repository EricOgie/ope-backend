package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strconv"
	"time"

	"github.com/EricOgie/ope-be/domain/models"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	mail "github.com/xhit/go-simple-mail/v2"
)

// Create an instance of SMTPClient that will be use for mailing. The function takes utils.Config struct
// and returns a *mail.SMTPClient.
// This way, we don't get to create multiple smtp connections, because we are creating just
// one instance that we are passing along the hexagonal routes to when and where it is needed
func GetEmailClient(env Config) *mail.SMTPClient {

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

// SendVerificationMail is callable when a verification mail mail needs to be sent
// It require an instance of responsedto.OneUserDtoWithOtp, and a token string to construct the required mailable
func SendVerificationMail(data responsedto.OneUserDtoWithOtp, token string) {
	env, loadErr := LoadConfig("*")
	if loadErr != nil {
		logger.Error("ENV_LOAD_ERR : " + loadErr.Error())
	}

	emailStruct := makeMailable(data, token)
	client := GetEmailClient(env)

	email := getEmail(emailStruct, env)
	err := email.Send(client)
	if err != nil {
		logger.Error(konstants.MAIL_DEL_ERR + err.Error())
		return
	}

}

// SendOTP is callable when OTP needs to be sent either for a 2FA fulfillment or
// It require an instance of responsedto.OneUserDtoWithOtp to construct the required mailable
func SendOTP(data responsedto.OneUserDtoWithOtp) {
	env, loadErr := LoadConfig("*")
	if loadErr != nil {
		logger.Error("ENV_LOAD_ERR : " + loadErr.Error())
	}

	emailStruct := makeOTPMailable(data)
	client := GetEmailClient(env)
	email := getOTPMail(emailStruct, env)
	err := email.Send(client)
	if err != nil {
		logger.Error(konstants.MAIL_DEL_ERR + err.Error())
		return
	}

}

func makeMailable(data responsedto.OneUserDtoWithOtp, token string) models.Emailable {
	redirURL := konstants.VERIFY_URL + "?k=" + token
	return models.Emailable{
		RecipientName:  data.FirstName,
		RecipientEmail: data.Email,
		Subject:        "Verify Your Ope Account",
		Body:           "Verify your Ope account",
		IsWithButton:   true,
		ButtonText:     "Verify Your Account",
		RedirectUrl:    redirURL,
		OTP:            strconv.Itoa(data.OTP),
	}
}

func makeOTPMailable(data responsedto.OneUserDtoWithOtp) models.Emailable {
	logger.Error(data.FirstName + "/" + strconv.Itoa(data.OTP))
	return models.Emailable{
		RecipientName:  data.FirstName,
		RecipientEmail: data.Email,
		Body:           konstants.MAIL_ACT_NOTICED,
		OTP:            strconv.Itoa(data.OTP),
	}
}

// ---------------- PRIVATE METHODS ---------------------//
func getEmail(emailStruct models.Emailable, env Config) *mail.Email {
	var body bytes.Buffer
	// mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; "
	body.Write([]byte(fmt.Sprintf("")))

	temp, err := template.ParseFiles("mailables/verification.html")
	if err != nil {
		logger.Error(konstants.MAIL_PARSE_ERR + err.Error())
	}

	temp.Execute(&body, emailStruct)
	// Create email
	email := mail.NewMSG()
	email.SetFrom(konstants.FROM_PREFIX + "<" + env.MailFromAddress + ">")
	email.AddTo(emailStruct.RecipientEmail)
	// email.AddCc(konstants.YAHOO)
	email.SetSubject(konstants.VERIFY_SUB)

	email.SetBody(mail.TextHTML, string(body.Bytes()))

	return email
}

func getOTPMail(emailStruct models.Emailable, env Config) *mail.Email {
	var body bytes.Buffer
	// mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; "
	body.Write([]byte(fmt.Sprintf("")))

	temp, err := template.ParseFiles("mailables/twofaemail.html")
	if err != nil {
		logger.Error(konstants.MAIL_PARSE_ERR + err.Error())
	}

	temp.Execute(&body, emailStruct)
	// Create email
	email := mail.NewMSG()
	email.SetFrom(konstants.FROM_PREFIX + "<" + env.MailFromAddress + ">")
	email.AddTo(emailStruct.RecipientEmail)
	// email.AddCc(konstants.YAHOO)
	email.SetSubject(konstants.VERIFY_SUB)

	email.SetBody(mail.TextHTML, string(body.Bytes()))

	return email
}
