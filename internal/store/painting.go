package store

import (
	"github.com/jinzhu/gorm"
	"github.com/sirtalin/democrart/internal/model"
	"github.com/sirtalin/democrart/internal/repository/painting"
)

type PaintingStore struct {
	db *gorm.DB
}

func NewPaintingStore(db *gorm.DB) *PaintingStore {
	return &PaintingStore{
		db: db,
	}
}

var _ painting.Store = &PaintingStore{}

// TODO: Validate Painting
// CreatePainting insert in the db the painting for the artist passed as argument
func (ps *PaintingStore) CreatePainting(artist *model.Artist, painting *model.Painting) error {
	tx := ps.db.Begin()
	err := ps.db.Model(artist).Association("Paintings").Append(painting).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, artMovement := range painting.Styles {
		err := tx.Where(&model.ArtMovement{Name: artMovement.Name}).First(&artMovement).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return err
		}
		if err := tx.Model(painting).Association("Styles").Append(artMovement).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, genre := range painting.Genres {
		err := tx.Where(&model.Genre{Name: genre.Name}).First(genre).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return err
		}
		if err := tx.Model(painting).Association("Genres").Append(genre).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, media := range painting.Medias {
		err := tx.Where(&model.Media{Name: media.Name}).First(&media).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return err
		}
		if err := tx.Model(painting).Association("Medias").Append(media).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// GetPainting returns the painting passed as parameter
func (ps *PaintingStore) GetPainting(artistName string, paintingName string) (*model.Painting, error) {
	var painting model.Painting
	ps.db.Model(&model.Painting{Name: artistName, Artist: &model.Artist{Name: paintingName}}).Preload("Image").Find(&painting)
	return &painting, nil
}

// CreateImage insert in the db the image passed as parameter
func (ps *PaintingStore) CreateImage(painting *model.Painting, image *model.Image) error {
	tx := ps.db.Begin()

	err := ps.db.Model(painting).Association("Images").Append(image).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
