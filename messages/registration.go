package messages

import (
	"github.com/studtool/common/queues"

	"github.com/studtool/emails-service/beans"
)

func (c *QueueClient) sendRegEmail(data []byte) {
	var regEmailData queues.RegistrationEmailData
	if err := c.parseMessageBody(data, &regEmailData); err != nil {
		beans.Logger().Error(err)
	} else {
		c.sendEmail(regEmailData.Email,
			c.regTemplate.Render(map[string]interface{}{
				"token": regEmailData.Token,
			}),
		)
	}
}
