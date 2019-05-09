package messages

import (
	"fmt"

	"github.com/mailru/easyjson"
	"github.com/streadway/amqp"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/queues"
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
		queues.RegistrationEmailsQueueName,
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

type MessageHandler func(data []byte)

func (c *QueueClient) Run() error {
	if err := c.receive(c.regQueue, c.sendRegEmail); err != nil {
		return err
	}
	return nil
}

func (c *QueueClient) receive(q amqp.Queue, handler MessageHandler) error {
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
			handler(d.Body)
		}
	}()

	return nil
}

func (c *QueueClient) parseMessageBody(data []byte, v easyjson.Unmarshaler) error {
	return easyjson.Unmarshal(data, v)
}

func (c *QueueClient) sendEmail(email, subject, text string) {
	if err := c.smtpClient.SendEmail(email, subject, text); err != nil {
		beans.Logger().Error(err)
	}
}
