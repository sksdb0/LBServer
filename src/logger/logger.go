package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var instance *Logger

type Logger struct {
	logger  *log.Logger
	logFile *os.File
}

func GetInstance() *Logger {
	if instance == nil {
		instance = new(Logger)
	}
	return instance
}

func (this *Logger) Init(name string) bool {
	logFile, err := os.Create(name)
	this.logFile = logFile
	if err != nil {
		log.Fatalln("open file error !")
		logFile.Close()
		return false
	}
	this.logger = log.New(logFile, "", log.Ltime)

	return true
}

func (this *Logger) Close() {
	this.logFile.Close()
}

func LOG(a ...interface{}) {
	GetInstance().logger.Print(a...)
}

func LOGLINE(a ...interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()

	newparam := make([]interface{}, 1)
	newparam[0] = funcname
	a = append(newparam, a...)

	GetInstance().logger.Println(a...)
}

func PRINTF(format string, a ...interface{}) {
	GetInstance().logger.Printf(format, a...)
}

func PRINTLINE(a ...interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()

	newparam := make([]interface{}, 1)
	newparam[0] = funcname
	a = append(newparam, a...)

	fmt.Println(a...)
	GetInstance().logger.Println(a...)
}
