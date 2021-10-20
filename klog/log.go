package klog

import (
	"log"
	"os"
	"runtime"
)

var (
	stdout = log.New(os.Stdout, "", log.LstdFlags)
	stderr = log.New(os.Stderr, "", log.LstdFlags)
)

func printLog(lg *log.Logger, x []interface{}, deep int) {
	_, file, line, _ := runtime.Caller(deep)
	// fname := runtime.FuncForPC(function).Name()
	if len(x) == 0 {
		lg.Printf("%s:%d\n", file, line)
	} else {
		lg.Printf("%s:%d\n- %+v\n", file, line, x)
	}
}

// Log ...
func Log(x ...interface{}) {
	printLog(stdout, x, 3)
}

// Error ...
func Error(x ...interface{}) {
	printLog(stderr, x, 3)
}

// Error1 ..
func Error1(x ...interface{}) {
	printLog(stderr, x, 2)
}
