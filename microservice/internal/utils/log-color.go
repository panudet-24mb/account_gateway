package utils

import (
	"fmt"
	"log"
)

//Info ...
var (
	Info    = Teal
	Warn    = Yellow
	Error   = Red
	Request = Purple
	Success = Green
)

//Black ...
var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

//Color ...
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

//InfoLog ...
func InfoLog(keys string, value interface{}) {
	log.Println(Info(keys, value))
}

//RequestLog ...
func RequestLog(keys string, value interface{}) {
	log.Println(Request(keys, value))
}

//WarningLog ...
func WarningLog(keys string, value interface{}) {
	log.Println(Warn(keys, value))
}

//ErrorLog ...
func ErrorLog(keys string, value interface{}) {
	log.Println(Error(keys, value))
}

//SuccessLog ...
func SuccessLog(keys string, value interface{}) {
	log.Println(Success(keys, value))
}
