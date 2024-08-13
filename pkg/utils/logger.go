package utils

import "log"

var logger = log.New(log.Writer(), "kevbeltrao/websocket: ", log.LstdFlags)

func SetLogger(customLogger *log.Logger) {
	if customLogger != nil {
		logger = customLogger
	}
}

func LogInfo(format string, v ...interface{}) {
	logger.Printf("INFO: "+format, v...)
}

func LogError(format string, v ...interface{}) {
	logger.Printf("ERROR: "+format, v...)
}
