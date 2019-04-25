package messages

import (
	"fmt"

	"github.com/streadway/amqp"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/utils"

	"github.com/studtool/emails-service/beans"
	"github.com/studtool/emails-service/config"
	"github.com/studtool/emails-service/emails"
	"github.com/studtool/emails-service/templates"
)

type QueueClient struct {
	connStr string

	ch   *amqp.Channel
	conn *amqp.Connection

	regQueue amqp.Queue

	smtpClient *emails.SmtpClient

	regTemplate *templates.RegistrationTemplate
}

func NewQueueClient(smtp *emails.SmtpClient,
	regTmp *templates.RegistrationTemplate) *QueueClient {

	return &QueueClient{
		connStr: fmt.Sprintf("amqp://%s:%s@%s:%s/",
			config.QueueUser.Value(), config.QueuePassword.Value(),
			config.QueueHost.Value(), config.QueuePort.Value(),
		),

		smtpClient:  smtp,
		regTemplate: regTmp,
	}
}

func (c *QueueClient) OpenConnection() error {
	var conn *amqp.Connection
	err := utils.WithRetry(func(n int) (err error) {
		if n > 0 {
			beans.Logger().Info(fmt.Sprintf("opening message queue connection. retry #%d", n))
		}
		conn, err = amqp.Dial(c.connStr)
		return err
	}, config.QueueConnNumRet.Value(), config.QueueConnRetItv.Value())
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	c.regQueue, err = ch.QueueDeclare(
		config.RegQueueName.Value(),
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	c.ch = ch
	c.conn = conn

	return nil
}

func (c *QueueClient) CloseConnection() error {
	if err := c.ch.Close(); err != nil {
		return err
	}
	return c.conn.Close()
}

type EmailRenderer func() string

func (c *QueueClient) Run() error {
	if err := c.receive(c.regQueue, c.renderRegEmail); err != nil {
		return err
	}
	return nil
}

func (c *QueueClient) receive(q amqp.Queue, r EmailRenderer) error {
	messages, err := c.ch.Consume(
		q.Name,
		consts.EmptyString,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range messages {
			c.sendEmail(string(d.Body), r())
		}
	}()

	return nil
}

func (c *QueueClient) sendEmail(email string, text string) {
	if err := c.smtpClient.SendEmail(email, text); err != nil {
		beans.Logger().Error(err)
	}
}

func (c *QueueClient) renderRegEmail() string {
	return c.regTemplate.Render(map[string]interface{}{})
}
