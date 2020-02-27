package log

import (
	"github.com/sirupsen/logrus"

	"github.com/kvault/gbgo/pkg/ipc"
)

// This wrapper allow me to print everything with the same format from anywhere
var logger *logrus.Logger

func init() {
	// TODO Some of this options must be set using viper
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...string) {
	ipc.LocChan(args...)
	logger.Debug(toInterface(args...))
}

// Info logs a message at level Info on the standard logger.
func Info(args ...string) {
	ipc.LocChan(args...)
	logger.Info(toInterface(args...))
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...string) {
	ipc.LocChan(args...)
	logger.Warn(toInterface(args...))
}

// Error logs a message at level Error on the standard logger.
func Error(args ...string) {
	ipc.LocChan(args...)
	logger.Error(toInterface(args...))
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...string) {
	ipc.LocChan(args...)
	logger.Fatal(toInterface(args...))
}

func toInterface(stringers ...string) []interface{} {
	newSlice := make([]interface{}, len(stringers))

	for i, s := range stringers {
		newSlice[i] = s
	}

	return newSlice
}
