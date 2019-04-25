package utils

import (
	"reflect"
	"runtime"
)

func NameOf(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
