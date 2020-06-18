package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// SetupLogs set the format for Logrus logs
func SetupLogs(logsLevel string, timestampFormat string, fullTimestamp bool) {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: timestampFormat,
		FullTimestamp:   fullTimestamp,
	})
	switch logsLevel {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
}

// SetupViper initialize viper with the config file in the path passed as parameter
func SetupViper(path string) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()

	if err != nil {
		logrus.Fatalf("Error while reading config file %s. %s", path, err)
	}
}
