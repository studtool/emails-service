package beans

import (
	"github.com/studtool/common/logs"
)

var (
	logger = logs.NewLogger()
)

func Logger() *logs.Logger {
	return logger
}
