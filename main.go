package main

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/dig"

	"github.com/studtool/common/utils"

	"github.com/studtool/emails-service/beans"
	"github.com/studtool/emails-service/emails"
	"github.com/studtool/emails-service/messages"
	"github.com/studtool/emails-service/templates"
)

func main() {
	c := dig.New()

	utils.AssertOk(c.Provide(templates.NewRegistrationTemplate))
	utils.AssertOk(c.Invoke(func(t *templates.RegistrationTemplate) {
		if err := t.Load(); err != nil {
			beans.Logger().Fatal(err)
		}
	}))

	utils.AssertOk(c.Provide(messages.NewQueueClient))
	utils.AssertOk(c.Invoke(func(c *messages.QueueClient) {
		if err := c.OpenConnection(); err != nil {
			beans.Logger().Fatal(err)
		}
	}))
	defer func() {
		utils.AssertOk(c.Invoke(func(c *messages.QueueClient) {
			if err := c.CloseConnection(); err != nil {
				beans.Logger().Fatal(err)
			}
		}))
	}()

	utils.AssertOk(c.Provide(emails.NewSmtpClient))

	var ch = make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)
	signal.Notify(ch, syscall.SIGINT)

	utils.AssertOk(c.Invoke(func(c *messages.QueueClient) {
		if err := c.Run(); err != nil {
			beans.Logger().Fatal(err)
		}
	}))

	<-ch
}
