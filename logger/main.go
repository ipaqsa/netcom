package logger

import (
	"log"
	"netcom/configurator"
	"os"
)

func NewLogger(typeLogger string) *log.Logger {
	f := os.Stderr
	if configurator.Debug != nil {
		if *configurator.Debug {
			f = os.Stdout
		}
	}
	if typeLogger == "ERROR" || typeLogger == "WARNING" {
		return log.New(f, typeLogger+": ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		return log.New(f, typeLogger+": ", log.Ldate|log.Ltime)
	}
}
