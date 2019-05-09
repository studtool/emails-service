package config

import (
	"time"

	"github.com/studtool/common/config"

	"github.com/studtool/emails-service/beans"
)

var (
	_ = func() *cconfig.FlagVar {
		f := cconfig.NewFlagDefault("STUDTOOL_EMAILS_SERVICE_SHOULD_LOG_ENV_VARS", false)
		if f.Value() {
			cconfig.SetLogger(beans.Logger())
		}
		return f
	}()

	SmtpHost     = cconfig.NewStringDefault("STUDTOOL_SMTP_SERVER_HOST", "127.0.0.1")
	SmtpPort     = cconfig.NewStringDefault("STUDTOOL_SMTP_SERVER_PORT", "25")
	SmtpUser     = cconfig.NewStringDefault("STUDTOOL_SMTP_SERVER_USER", "user")
	SmtpPassword = cconfig.NewStringDefault("STUDTOOL_SMTP_SERVER_PASSWORD", "password")
	SmtpSSL      = cconfig.NewFlagDefault("STUDTOOL_SMTP_SERVER_SSL", true)

	QueueHost     = cconfig.NewStringDefault("STUDTOOL_EMAILS_QUEUE_HOST", "127.0.0.1")
	QueuePort     = cconfig.NewStringDefault("STUDTOOL_EMAILS_QUEUE_PORT", "5672")
	QueueUser     = cconfig.NewStringDefault("STUDTOOL_EMAILS_QUEUE_USER", "user")
	QueuePassword = cconfig.NewStringDefault("STUDTOOL_EMAILS_QUEUE_PASSWORD", "password")

	QueueConnNumRet = cconfig.NewIntDefault("STUDTOOL_EMAILS_QUEUE_CONNECTION_NUM_RETRIES", 10)
	QueueConnRetItv = cconfig.NewTimeSecsDefault("STUDTOOL_EMAILS_QUEUE_CONNECTION_RETRY_INTERVAL", 2*time.Second)

	RegEmailTemplatePath = cconfig.NewStringDefault("STUDTOOL_REGISTRATION_EMAIL_TEMPLATE_PATH", "./template.txt")
)
