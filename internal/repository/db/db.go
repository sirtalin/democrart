package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Required by the ORM
	gormlog "github.com/onrik/logrus/gorm"
	"github.com/sirtalin/democrart/internal/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// New returns a gorm database connection
func New() *gorm.DB {
	var driver string = viper.GetString("database_driver")
	var user string = viper.GetString("database_user")
	var password string = viper.GetString("database_password")
	var port string = viper.GetString("database_port")
	var host string = viper.GetString("database_host")
	var sslMode string = viper.GetString("database_ssl")
	var databaseName string = viper.GetString("database_name")
	var dbConn string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, databaseName, sslMode)

	var maxIdleCons int = viper.GetInt("database_max_idle_conns")
	var maxOpenConns int = viper.GetInt("database_max_open_conns")

	db, err := gorm.Open(driver, dbConn)
	if err != nil {
		logrus.Fatalf("Error connecting to database %s. %s.", databaseName, err)
	}

	db.SetLogger(gormlog.New(logrus.StandardLogger()))
	db.LogMode(true)

	db.DB().SetMaxIdleConns(maxIdleCons)
	db.DB().SetMaxOpenConns(maxOpenConns)

	return db
}

//TODO: check errors
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.ArtMovement{},
		&model.Nationality{},
		&model.PaintingSchool{},
		&model.Artist{},
		&model.Media{},
		&model.Genre{},
	)
	db.AutoMigrate(&model.ArtistsSchool{}).AddForeignKey("artist_id", "artists(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.ArtistsSchool{}).AddForeignKey("painting_school_id", "painting_schools(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.ArtistsNationality{}).AddForeignKey("artist_id", "artists(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.ArtistsNationality{}).AddForeignKey("nationality_id", "nationalities(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.ArtistsMovement{}).AddForeignKey("artist_id", "artists(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.ArtistsMovement{}).AddForeignKey("art_movement_id", "art_movements(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.Painting{}).AddForeignKey("artist_id", "artists(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.Image{}).AddForeignKey("painting_id", "paintings(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.PaintingsStyle{}).AddForeignKey("painting_id", "paintings(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.PaintingsStyle{}).AddForeignKey("art_movement_id", "art_movements(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.PaintingsGenre{}).AddForeignKey("painting_id", "paintings(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.PaintingsGenre{}).AddForeignKey("genre_id", "genres(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.PaintingsMedias{}).AddForeignKey("painting_id", "paintings(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.PaintingsMedias{}).AddForeignKey("media_id", "media(id)", "RESTRICT", "RESTRICT")
}
