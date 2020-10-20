package main

import (
	"github.com/sirtalin/democrart/internal/handler"
	"github.com/sirtalin/democrart/internal/repository/db"
	"github.com/sirtalin/democrart/internal/scraper"
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
	var scrapartHandler *handler.Handler = handler.NewHandler(artistStore, paintingStore)

	var artMovementURLs []string = scraper.GetArtMovementURLs()

	for _, artMovementURL := range artMovementURLs {
		scraper.GetArtMovement(artMovementURL, scrapartHandler)
	}
}
