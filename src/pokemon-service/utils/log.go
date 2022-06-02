package utils

import (
	"log"
	"os"
	"sync"
)

type logUtils struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	panicLogger   *log.Logger
}

var lock = &sync.Mutex{}
var singleInstance *logUtils

func getInstance() *logUtils {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &logUtils{}
			file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			singleInstance.errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
			singleInstance.infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
			singleInstance.panicLogger = log.New(file, "PANIC: ", log.Ldate|log.Ltime|log.Lshortfile)
			singleInstance.warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		}
	}

	return singleInstance
}
