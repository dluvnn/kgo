package kgo

import (
	"encoding/json"
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
	printLog(stdout, x, 2)
}

// Error ...
func Error(x ...interface{}) {
	printLog(stderr, x, 2)
}

func PrintJSON(x interface{}) {
	data, err := json.MarshalIndent(x, "", "    ")
	if err != nil {
		Error(err)
		return
	}
	println(string(data))
}
