package cconfig

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/studtool/common/consts"
)

type TimeVar struct {
	value time.Duration
}

func NewTime(name string) *TimeVar {
	return parseTimeSecs(name, consts.AnyTime, true)
}

func NewTimeDefault(name string, defVal time.Duration) *TimeVar {
	return parseTimeSecs(name, defVal, false)
}

func (v *TimeVar) Value() time.Duration {
	return v.value
}

func parseTimeSecs(name string, defVal time.Duration, isRequired bool) *TimeVar {
	var t time.Duration

	v := os.Getenv(name)
	if v == consts.EmptyString {
		if isRequired {
			panicNotSet(name)
		} else {
			t = defVal
		}
	} else {
		m := v[len(v)-1]

		switch m {
		case 's':
			t = time.Second
		case 'm':
			t = time.Minute
		case 'h':
			t = time.Hour
		default:
			panicInvalidFormat(name, "[INTEGER]s")
		}

		tVal, err := strconv.Atoi(v[:len(v)-1])
		if err != nil {
			panic(err)
		}

		t *= time.Duration(tVal)
	}

	if logger != nil {
		logger.Info(fmt.Sprintf("%s=%v", name, t))
	}

	return &TimeVar{
		value: t,
	}
}
