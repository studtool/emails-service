package messages

import (
	"fmt"

	"github.com/mailru/easyjson"
	"github.com/streadway/amqp"
	"go.uber.org/dig"

	"github.com/studtool/common/consts"
	"github.com/studtool/common/queues"
	"github.com/studtool/common/utils"

	"github.com/studtool/emails-service/beans"
	"github.com/studtool/emails-service/config"
	"github.com/studtool/emails-service/emails"
	"github.com/studtool/emails-service/templates"
)

type QueueClient struct {
	connStr    string
	connection *amqp.Connection

	channel *amqp.Channel

	smtpClient *emails.SmtpClient

	regEmailsQueue   amqp.Queue
	regEmailTemplate *templates.RegistrationTemplate
}

type ClientParams struct {
	dig.In

	SmtpClient       *emails.SmtpClient
	RegEmailTemplate *templates.RegistrationTemplate
}

func NewQueueClient(params ClientParams) *QueueClient {
	return &QueueClient{
		connStr: fmt.Sprintf("amqp://%s:%s@%s:%s/",
			config.MqUser.Value(), config.MqPassword.Value(),
			config.MqHost.Value(), config.MqPort.Value(),
		),
		smtpClient:       params.SmtpClient,
		regEmailTemplate: params.RegEmailTemplate,
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
	}, config.MqConnNumRet.Value(), config.MqConnRetItv.Value())
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	c.regEmailsQueue, err = ch.QueueDeclare(
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

	c.channel = ch
	c.connection = conn

	return nil
}

func (c *QueueClient) CloseConnection() error {
	if err := c.channel.Close(); err != nil {
		return err
	}
	return c.connection.Close()
}

type MessageHandler func(data []byte)

func (c *QueueClient) Run() error {
	if err := c.receiveRegEmailMessages(); err != nil {
		return err
	}
	return nil
}

func (c *QueueClient) receiveRegEmailMessages() error {
	messages, err := c.channel.Consume(
		c.regEmailsQueue.Name,
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
			c.sendRegEmail(d.Body)
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

func (c *QueueClient) handleErr(err error) {
	beans.Logger().Error(err)
}
