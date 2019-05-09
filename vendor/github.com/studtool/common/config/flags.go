package config

import (
	"fmt"
	"os"

	"github.com/studtool/common/consts"
)

var (
	//nolint:gochecknoglobals
	validTrueValues = []string{"TRUE", "true", "True", "1"}

	//nolint:gochecknoglobals
	validFalseValues = []string{"FALSE", "false", "False", "0"}

	//nolint:gochecknoglobals
	flagValues = func() map[string]bool {
		m := make(map[string]bool)
		for _, t := range validTrueValues {
			m[t] = true
		}
		for _, f := range validFalseValues {
			m[f] = false
		}
		return m
	}()
)

type FlagVar struct {
	value bool
}

func NewFlag(name string) *FlagVar {
	return parseFlag(name, consts.AnyBool, true)
}

func NewFlagDefault(name string, defVal bool) *FlagVar {
	return parseFlag(name, defVal, false)
}

func (v *FlagVar) Value() bool {
	return v.value
}

func parseFlag(name string, defVal bool, isRequired bool) *FlagVar {
	var f bool

	v := os.Getenv(name)
	if v == consts.EmptyString {
		if isRequired {
			panicNotSet(name)
		} else {
			f = defVal
		}
	} else {
		var ok bool
		if f, ok = flagValues[v]; !ok {
			panicInvalidFormat(name, fmt.Sprintf("one of %v", flagValues))
		}
	}

	if logger != nil {
		logger.Info(fmt.Sprintf("%s=%v", name, f))
	}

	return &FlagVar{
		value: f,
	}
}
