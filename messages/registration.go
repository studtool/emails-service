package messages

import (
	"github.com/studtool/common/queues"

	"github.com/studtool/emails-service/beans"
)

func (c *QueueClient) sendRegEmail(body []byte) {
	data := &queues.RegistrationEmailData{}
	if err := c.parseMessageBody(body, data); err != nil {
		beans.Logger().Error(err)
	} else {
		c.sendEmail(
			data.Email, "Registration",
			c.regEmailTemplate.Render(map[string]interface{}{
				"token": data.Token,
			}),
		)
	}
}
