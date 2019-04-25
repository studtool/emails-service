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

type MQ struct {
	connStr string

	ch   *amqp.Channel
	conn *amqp.Connection

	regQueue    amqp.Queue
	regTemplate *templates.RegistrationTemplate

	smtpClient *emails.SmtpClient
}

func NewQueue() *MQ {
	return &MQ{
		connStr: fmt.Sprintf("amqp://%s:%s@%s:%s/",
			config.QueueUser.Value(), config.QueuePassword.Value(),
			config.QueueHost.Value(), config.QueuePort.Value(),
		),
	}
}

func (mq *MQ) OpenConnection() error {
	var conn *amqp.Connection
	err := utils.WithRetry(func(n int) (err error) {
		if n > 0 {
			beans.Logger().Info(fmt.Sprintf("opening message queue connection. retry #%d", n))
		}
		conn, err = amqp.Dial(mq.connStr)
		return err
	}, config.QueueConnNumRet.Value(), config.QueueConnRetItv.Value())
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	mq.regQueue, err = ch.QueueDeclare(
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

	mq.ch = ch
	mq.conn = conn

	return nil
}

func (mq *MQ) CloseConnection() error {
	if err := mq.ch.Close(); err != nil {
		return err
	}
	return mq.conn.Close()
}

type EmailRenderer func() string

func (mq *MQ) Run() error {
	if err := mq.receive(mq.regQueue, mq.renderRegEmail); err != nil {
		return err
	}
	return nil
}

func (mq *MQ) receive(q amqp.Queue, r EmailRenderer) error {
	messages, err := mq.ch.Consume(
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
			mq.sendEmail(string(d.Body), r())
		}
	}()

	return nil
}

func (mq *MQ) sendEmail(email string, text string) {
	if err := mq.smtpClient.SendEmail(email, text); err != nil {
		beans.Logger().Error(err)
	}
}

func (mq *MQ) renderRegEmail() string {
	return mq.regTemplate.Render(map[string]interface{}{})
}
