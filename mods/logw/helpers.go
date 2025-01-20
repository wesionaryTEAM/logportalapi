package logw

import (
	"fmt"
	"runtime"
)

func GetCaller(skip int) string {
	caller := ""

	_, file, line, ok := runtime.Caller(skip)

	if ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	}

	return caller
}
