package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/studtool/common/consts"
)

type TimeSecsVar struct {
	value time.Duration
}

func NewTimeSecs(name string) *TimeSecsVar {
	return parseTimeSecs(name, consts.AnyTime, true)
}

func NewTimeSecsDefault(name string, defVal time.Duration) *TimeSecsVar {
	return parseTimeSecs(name, defVal, false)
}

func (v *TimeSecsVar) Value() time.Duration {
	return v.value
}

func parseTimeSecs(name string, defVal time.Duration, isRequired bool) *TimeSecsVar {
	var t time.Duration

	v := os.Getenv(name)
	if v == consts.EmptyString {
		if isRequired {
			panicNotSet(name)
		} else {
			t = defVal
		}
	} else {
		if v[len(v)-1] == 's' {
			v = v[:len(v)-1]

			tVal, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			t = time.Duration(tVal) * time.Second
		} else {
			panicInvalidFormat(name, "[INTEGER]s")
		}
	}

	if logger != nil {
		logger.Info(fmt.Sprintf("%s=%v", name, t))
	}

	return &TimeSecsVar{
		value: t,
	}
}
