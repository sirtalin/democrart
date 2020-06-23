package store

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sirtalin/democrart/internal/model"
	"github.com/sirtalin/democrart/internal/repository/painting"
	"github.com/sirtalin/democrart/pkg/utils"
	"github.com/sirupsen/logrus"
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

func (ps *PaintingStore) GetPaintings(painting *model.Painting) ([]*model.Artist, error) {
	var artists []*model.Artist
	var ids []uint
	var bids []uint
	var firstFilter bool = true

	var tx *gorm.DB = ps.db

	var genre model.Genre
	var paintingsByGenre []*model.Painting

	var style model.ArtMovement
	var paintingsByStyle []*model.Painting

	var media model.Media
	var paintingsByMedia []*model.Painting

	var paintingsByName []*model.Painting

	if len(painting.Genres) == 1 {
		err := ps.db.Where(&model.Genre{Name: painting.Genres[0].Name}).First(&genre).Error

		if err != nil {
			return nil, err
		}

		err = ps.db.Model(&genre).Related(&paintingsByGenre, "Paintings").Error

		if !firstFilter {
			bids = ids
			ids = nil
		}

		for _, p := range paintingsByGenre {
			if firstFilter || utils.UintInSlice(p.ID, bids) {
				ids = append(ids, p.ID)
			}
		}

		firstFilter = false
	}

	if len(painting.Styles) == 1 {
		err := ps.db.Where(&model.ArtMovement{Name: painting.Styles[0].Name}).First(&style).Error

		if err != nil {
			return nil, err
		}

		err = ps.db.Model(&style).Related(&paintingsByStyle, "Paintings").Error

		if !firstFilter {
			bids = ids
			ids = nil
		}

		for _, p := range paintingsByStyle {
			if firstFilter || utils.UintInSlice(p.ID, bids) {
				ids = append(ids, p.ID)
			}
		}

		firstFilter = false
	}

	if len(painting.Medias) == 1 {
		err := ps.db.Where(&model.Media{Name: painting.Medias[0].Name}).First(&media).Error

		if err != nil {
			return nil, err
		}

		err = ps.db.Model(&media).Related(&paintingsByMedia, "Paintings").Error

		if !firstFilter {
			bids = ids
			ids = nil
		}

		for _, p := range paintingsByMedia {
			if firstFilter || utils.UintInSlice(p.ID, bids) {
				ids = append(ids, p.ID)
			}
		}

		firstFilter = false
	}

	if len(painting.Name) >= 3 {
		err := ps.db.Where("name LIKE ?", fmt.Sprintf("%s%s%s", "%", painting.Name, "%")).Find(&paintingsByName).Error

		if err != nil {
			return nil, err
		}

		if !firstFilter {
			bids = ids
			ids = nil
		}

		for _, p := range paintingsByName {
			if firstFilter || utils.UintInSlice(p.ID, bids) {
				ids = append(ids, p.ID)
			}
		}

		firstFilter = false
	}

	err := tx.Preload("Paintings", "id in (?)", ids).Preload("Paintings.Genres").Preload("Paintings.Styles").Preload("Paintings.Medias").Find(&artists).Error

	if err != nil {
		logrus.Error(err)
		if gorm.IsRecordNotFoundError(err) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return artists, nil
}
