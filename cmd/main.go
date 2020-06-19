package main

import (
	"github.com/sirtalin/democrart/internal/repository/db"
	"github.com/sirtalin/democrart/pkg/utils"

	"github.com/spf13/viper"
)

func init() {
	utils.SetupViper(".env")
	utils.SetupLogs(viper.GetString("logs_level"), viper.GetString("timestamp_format"), viper.GetBool("full_timestamp"))
}

func main() {
	dbConn := db.New()
	db.AutoMigrate(dbConn)
	defer dbConn.Close()
}
