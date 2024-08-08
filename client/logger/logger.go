package logger

import "log"

func Error(msg string, err error) {
	log.Fatalf("%s %v", msg, err)
}
