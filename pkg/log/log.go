package log

import (
	"fmt"
	"strings"
	"time"
)

func log(LogLevel string, message string) {
	fmt.Println(fmt.Sprintf("[%s] %s %s", strings.ToUpper(LogLevel), time.Now().Format("2006-01-02 15:04:05"), message))
}

func Error(message string, args ...any) {
	log("error", fmt.Sprintf(message, args...))
}

func Info(message string, args ...any) {
	log("info", fmt.Sprintf(message, args...))
}

func Debug(message string, args ...any) {
	log("debug", fmt.Sprintf(message, args...))
}

func Panic(err error) {
	if err != nil {
		Error(err.Error())
		panic(err)
	}
}
