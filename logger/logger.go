package logger

import (
	"fmt"
	"lemon-robot-cli/utils/date"
)

func Debug(msg string) {
	log(msg, 0)
}

func Warn(msg string) {
	log(msg, 1)
}

func Error(msg string) {
	log(msg, 2)
}

func Info(msg string) {
	log(msg, 3)
}

var logTypeList = [] string{"DEBG", "WARN", "ERRO", "INFO"}
var logColorList = [] int{0, 33, 31, 34}

func log(msg string, logType int) {
	fmt.Printf("%c[0;0;%dm[%s %s] %s%c[0m\n", 0x1B, logColorList[logType], logTypeList[logType], date.GetCurrentTimeFormatedStr(), msg, 0x1B)
}
