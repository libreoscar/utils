package assert

import (
	"fmt"
	"runtime"
)

// The goal of this package is to alleviate the pain of repetitive code in defensive
// programming. Do not abuse it.

// returns file_name, line_number
func getCallSite(lvl int) (string, int) {
	_, file, line, ok := runtime.Caller(lvl)
	if !ok {
		return "", 0
	} else {
		return file, line
	}
}

func True(v bool) {
	if !v {
		file, num := getCallSite(2)
		panic(fmt.Errorf("assertion failed at %s:%d", file, num))
	}
}

func False(v bool) {
	if v {
		file, num := getCallSite(2)
		panic(fmt.Errorf("assertion failed at %s:%d", file, num))
	}
}

func ErrorIsNil(e error) {
	if e != nil {
		panic(e)
	}
}
