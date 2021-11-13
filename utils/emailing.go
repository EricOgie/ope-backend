package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"

	"github.com/EricOgie/ope-be/domain/models"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	mail "github.com/xhit/go-simple-mail/v2"
)

// SendVerificationMail is callable when a verification mail mail needs to be sent
// It require an instance of responsedto.OneUserDtoWithOtp, and a token string to construct the required mailable
func SendVerificationMail(data responsedto.OneUserDtoWithOtp, token string) {
	executeMailing(data, token, konstants.MAIL_PURPOSE_VERIFY, konstants.MAIL_BTN_VET, konstants.MAIL_VET_PATH)
}

// SendOTP is callable when OTP needs to be sent either for a 2FA fulfillment or
// It require an instance of responsedto.OneUserDtoWithOtp to construct the required mailable
func SendOTP(data responsedto.OneUserDtoWithOtp) {
	executeMailing(data, "", konstants.MAIL_PURPOSE_OTP, "", konstants.MAIL_OTP_PATH)
}

// SendRequestMail is callable when a password change is to be requested.
// It require an instance of responsedto.OneUserDtoWithOtp to construct the required mailable
func SendRequestMail(data responsedto.OneUserDtoWithOtp) {
	executeMailing(data, "", konstants.MAIL_PURPOSE_REQ, konstants.MAIL_BTN_PWORD, konstants.MAIL_VET_PATH)
}

// ---------------- PRIVATE METHODS ---------------------//

func executeMailing(data responsedto.OneUserDtoWithOtp, token string,
	purpose string, btnTxt string, TemplatePATH string) {
	env := LoadConfig(".")
	emailStruct := makeMailable(data, token, purpose, btnTxt)
	client := getEmailClient(env)
	email := makeEmail(emailStruct, env, "mailables/vet.html")
	err := email.Send(client)
	if err != nil {
		logger.Error(konstants.MAIL_DEL_ERR + err.Error())
		return
	}
}

func makeEmail(emailStruct models.Emailable, env Config, temPath string) *mail.Email {
	var body bytes.Buffer
	// mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; "
	body.Write([]byte(fmt.Sprintf("")))

	temp, err := template.ParseFiles(temPath)
	if err != nil {
		logger.Error(konstants.MAIL_PARSE_ERR + err.Error())
	}
	//mashall emailstruct into email body
	temp.Execute(&body, emailStruct)
	// Create email
	email := mail.NewMSG()
	email.SetFrom(konstants.FROM_PREFIX + "<" + env.MailFromAddress + ">")
	email.AddTo(emailStruct.RecipientEmail)
	email.SetSubject(emailStruct.Subject)
	email.SetBody(mail.TextHTML, string(body.Bytes()))

	return email
}

func makeMailable(data responsedto.OneUserDtoWithOtp, token string,
	purpose string, btnTxt string) models.Emailable {
	redirURL := ""
	subject := ""
	body := ""
	tail := ""
	caption := ""

	if purpose == "verification" {
		redirURL += konstants.ROOT_ADD + "verify-account?k=" + token
		subject += konstants.SUBJECT_VERIFY_ACC
		body += konstants.MAIL_BODY_VERIFY
		tail += konstants.MAIL_TAIL_VERIFY
		caption += konstants.CAPTION_WELCOME
	} else if purpose == "OTP" {
		redirURL += ""
		subject += konstants.SUBJECT_OTP
		body += konstants.MAIL_BODY_OTP
		tail += konstants.MAIL_TAIL_ACT_NOTICED
		caption += konstants.CAPTION_HELLO
	} else {
		redirURL += konstants.ROOT_ADD
		subject += konstants.SUBJECT_PASSWORD_CHANGE
		body += konstants.MAIL_BODY_PASSWORD_RESET
		tail += konstants.MAIL_TAIL_PASSWORD_REQ
		caption += konstants.CAPTION_HELLO
	}

	return models.Emailable{
		RecipientName:  data.FirstName,
		RecipientEmail: data.Email,
		Subject:        subject,
		Caption:        caption,
		Body:           body,
		Tail:           tail,
		IsWithButton:   true,
		ButtonText:     btnTxt,
		RedirectUrl:    redirURL,
		OTP:            strconv.Itoa(data.OTP),
	}
}
