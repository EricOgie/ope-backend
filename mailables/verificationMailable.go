package mailables

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/utils"
	mail "github.com/xhit/go-simple-mail/v2"
)

// SendVerification email takes mail.SMTPClient, utils.Config, and models.Emailable
// It will construct the verification mail and proceed to send.
// It will return nill on success mail sent, otherwise return the error
func SendVerificationMail(client mail.SMTPClient, env utils.Config, emailStruct models.Emailable) *ericerrors.EricError {

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Verify Your Ope Account \n%s\n\n", mimeHeaders)))

	temp, err := template.ParseFiles("verification.html")
	if err != nil {
		logger.Error(konstants.MAIL_PARSE_ERR + err.Error())
	}

	temp.Execute(&body, emailStruct)

	// Create email
	email := mail.NewMSG()
	email.SetFrom(konstants.FROM_PREFIX + env.MailFromAddress)
	email.AddTo(emailStruct.RecipientEmail)
	email.AddCc(konstants.YAHOO)
	email.SetSubject(konstants.VERIFY_SUB)

	email.SetBody(mail.TextHTML, string(body.Bytes()))

	err = email.Send(&client)
	if err != nil {
		logger.Error(konstants.MAIL_DEL_ERR)
		return ericerrors.NewError(44, err.Error())
	} else {
		return nil
	}

}
