package emails

import (
	"crypto/tls"
	"fmt"
	"net/mail"
	"net/smtp"

	"github.com/studtool/common/consts"

	"github.com/studtool/emails-service/config"
)

type SmtpClient struct {
	sendFunc SendFunc
}

func NewSmtpClient() *SmtpClient {
	c := &SmtpClient{}
	if config.SmtpSSL.Value() {
		c.sendFunc = c.SendEmailTLS
	} else {
		panic("not implemented") //TODO
	}
	return c
}

type SendFunc func(email, subject, text string) error

func (c *SmtpClient) SendEmailTLS(email, subject, text string) (err error) {
	from := mail.Address{
		Name:    consts.EmptyString,
		Address: config.SmtpUser.Value(),
	}
	to := mail.Address{
		Name:    consts.EmptyString,
		Address: email,
	}

	subj := subject
	body := text

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	message := consts.EmptyString
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += fmt.Sprintf("\r\n%s", body)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         config.SmtpHost.Value(),
	}

	addr := fmt.Sprintf("%s:%s",
		config.SmtpHost.Value(), config.SmtpPort.Value())

	var conn *tls.Conn
	conn, err = tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return err
	}
	defer func() {
		err = conn.Close()
	}()

	var client *smtp.Client
	client, err = smtp.NewClient(conn, config.SmtpHost.Value())
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth(
		consts.EmptyString,
		config.SmtpUser.Value(),
		config.SmtpPassword.Value(),
		config.SmtpHost.Value(),
	)
	if err = client.Auth(auth); err != nil {
		return err
	}

	if err = client.Mail(from.Address); err != nil {
		return err
	}
	if err = client.Rcpt(to.Address); err != nil {
		return err
	}

	w, err := client.Data()
	if err != nil {
		return err
	}
	defer func() {
		err = w.Close()
	}()

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	return client.Quit()
}
