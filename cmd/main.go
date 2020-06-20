package main

import (
	"github.com/sirtalin/democrart/internal/handler"
	"github.com/sirtalin/democrart/internal/repository/db"
	"github.com/sirtalin/democrart/internal/route"
	"github.com/sirtalin/democrart/internal/store"
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

	var artistStore *store.ArtistStore = store.NewArtistStore(dbConn)
	var paintingStore *store.PaintingStore = store.NewPaintingStore(dbConn)
	var democrartHandler *handler.Handler = handler.NewHandler(artistStore, paintingStore)

	router := route.New()
	api := router.Group("/api")

	democrartHandler.Register(api)
	router.Logger.Fatal(router.Start(":3000"))
}
