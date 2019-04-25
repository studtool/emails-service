package config

import (
	"fmt"
	"os"

	"github.com/studtool/common/consts"
)

type StringVar struct {
	value string
}

func NewString(name string) *StringVar {
	return parseString(name, consts.AnyString, false)
}

func NewStringDefault(name string, defVal string) *StringVar {
	return parseString(name, defVal, false)
}

func (v *StringVar) Value() string {
	return v.value
}

func parseString(name string, defVal string, isRequired bool) *StringVar {
	v := os.Getenv(name)
	if v == consts.EmptyString {
		if isRequired {
			panic(fmt.Sprintf("config: %s is required", name))
		} else {
			v = defVal
		}
	}

	if logger != nil {
		logger.Info(fmt.Sprintf("%s=%s", name, v))
	}

	return &StringVar{
		value: v,
	}
}
