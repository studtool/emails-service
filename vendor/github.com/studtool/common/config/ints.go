package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/studtool/common/consts"
)

type IntVar struct {
	value int
}

func NewInt(name string) *IntVar {
	return parseInt(name, consts.AnyInt, true)
}

func NewIntDefault(name string, defVal int) *IntVar {
	return parseInt(name, defVal, false)
}

func (v *IntVar) Value() int {
	return v.value
}

func parseInt(name string, defVal int, isRequired bool) *IntVar {
	var t int

	v := os.Getenv(name)
	if v == consts.EmptyString {
		if isRequired {
			panicNotSet(name)
		} else {
			t = defVal
		}
	} else {
		var err error
		if t, err = strconv.Atoi(v); err != nil {
			panicInvalidFormat(name, "[INTEGER]")
		}
	}

	if logger != nil {
		logger.Info(fmt.Sprintf("%s=%v", name, t))
	}

	return &IntVar{
		value: t,
	}
}
