package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Creates and returns a new logrus instance
func NewLogger() *logrus.Logger {
	// Create a new instance of the logger. You can have any number of instances.
	var log = logrus.New()

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(logrus.InfoLevel)

	return log
}
