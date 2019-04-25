package config

import (
	"time"

	"github.com/studtool/common/config"

	"github.com/studtool/emails-service/beans"
)

var (
	_ = func() *config.FlagVar {
		f := config.NewFlagDefault("STUDTOOL_EMAILS_SERVICE_SHOULD_LOG_ENV_VARS", false)
		if f.Value() {
			config.SetLogger(beans.Logger())
		}
		return f
	}()

	SmtpHost     = config.NewStringDefault("STUDTOOL_SMTP_SERVER_HOST", "127.0.0.1")
	SmtpPort     = config.NewStringDefault("STUDTOOL_SMTP_SERVER_PORT", "25")
	SmtpUser     = config.NewStringDefault("STUDTOOL_SMTP_SERVER_USER", "user")
	SmtpPassword = config.NewStringDefault("STUDTOOL_SMTP_SERVER_PASSWORD", "password")

	QueueHost     = config.NewStringDefault("STUDTOOL_EMAILS_QUEUE_HOST", "127.0.0.1")
	QueuePort     = config.NewStringDefault("STUDTOOL_EMAILS_QUEUE_PORT", "5672")
	QueueUser     = config.NewStringDefault("STUDTOOL_EMAILS_QUEUE_USER", "user")
	QueuePassword = config.NewStringDefault("STUDTOOL_EMAILS_QUEUE_PASSWORD", "password")

	QueueConnNumRet = config.NewIntDefault("STUDTOOL_EMAILS_QUEUE_CONNECTION_NUM_RETRIES", 10)
	QueueConnRetItv = config.NewTimeSecsDefault("STUDTOOL_EMAILS_QUEUE_CONNECTION_RETRY_INTERVAL", 2*time.Second)

	RegEmailTemplatePath = config.NewStringDefault("STUDTOOL_REGISTRATION_EMAIL_TEMPLATE_PATH", "./template.txt")
)
