package emails

import (
	"fmt"
	"net/smtp"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/errs"

	"github.com/studtool/emails-service/config"
)

type SmtpClient struct{}

func NewSmtpClient() *SmtpClient {
	return &SmtpClient{}
}

func (c *SmtpClient) SendEmail(email string, text string) *errs.Error {
	auth := smtp.PlainAuth(
		consts.EmptyString,
		config.SmtpUser.Value(),
		config.SmtpPassword.Value(),
		config.SmtpHost.Value(),
	)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s",
			config.SmtpHost.Value(), config.SmtpPort.Value(),
		),
		auth,
		config.SmtpUser.Value(),
		[]string{email},
		[]byte(text),
	)
	if err != nil {
		return errs.New(err)
	}

	return nil
}
