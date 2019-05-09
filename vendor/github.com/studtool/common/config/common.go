package config

import (
	"fmt"

	"github.com/studtool/common/logs"
)

var (
	//nolint:gochecknoglobals
	logger *logs.Logger
)

func SetLogger(log *logs.Logger) {
	logger = log
}

func panicNotSet(name string) {
	panic(fmt.Sprintf("config: %s is required", name))
}

func panicInvalidFormat(name string, pattern string) {
	panic(fmt.Sprintf("config: %s format error; pattern - %s", name, pattern))
}
