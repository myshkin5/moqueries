package logs

import (
	"log"
	"os"
)

var (
	debug      = true
	warnLogger *log.Logger
	errLogger  *log.Logger
)

func init() {
	errLogger = log.New(os.Stderr, "ERROR: ", log.LstdFlags)
	warnLogger = log.New(os.Stderr, "WARNING: ", log.LstdFlags)
	log.SetPrefix("DEBUG: ")
}

// Init initializes the loggers
func Init(dbg bool) {
	debug = dbg
}

// Debugf will output a debugging log if debug logs are configured
func Debugf(format string, args ...interface{}) {
	if debug {
		log.Printf(format, args...)
	}
}

// Warn logs a warning message
func Warn(msg string) {
	warnLogger.Printf(msg + ": %v")
}

// Error logs a error message with an error
func Error(msg string, err error) {
	errLogger.Printf(msg+": %v", err)
}

// Errorf logs a formatted message
func Errorf(format string, args ...interface{}) {
	errLogger.Printf(format, args...)
}

// Panic logs a message with an error then panics
func Panic(msg string, err error) {
	log.Panicf(msg+": %v", err)
}

// Panicf logs a formatted message then panics
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}
