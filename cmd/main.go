package main

import (
	"github.com/sirtalin/democrart/pkg/utils"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func init() {
	utils.SetupViper(".env")
	utils.SetupLogs(viper.GetString("logs_level"), viper.GetString("timestamp_format"), viper.GetBool("full_timestamp"))
}

func main() {
	logrus.Info("Hello World")
}
