package store

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sirtalin/democrart/internal/model"
)

type ArtistStore struct {
	db *gorm.DB
}

func NewArtistStore(db *gorm.DB) *ArtistStore {
	return &ArtistStore{
		db: db,
	}
}

// GetArtist returns the artist model of the artist with that name
func (as *ArtistStore) GetArtist(artistName string) (*model.Artist, error) {
	var artist model.Artist
	err := as.db.Where(&model.Artist{Name: artistName}).First(&artist).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &artist, err
}

// TODO: Validate Artist
// CreateArtist inserts in the db the artist passed as argument or returns an error
func (as *ArtistStore) CreateArtist(artist *model.Artist) error {
	tx := as.db.Begin()
	if err := tx.Create(artist).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, nationality := range artist.Nationalities {
		err := tx.Where(&model.Nationality{Demonym: nationality.Demonym}).First(&nationality).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return err
		}
		if err := tx.Model(artist).Association("Nationalities").Append(nationality).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, paintingSchool := range artist.PaintingSchools {
		err := tx.Where(&model.PaintingSchool{Name: paintingSchool.Name}).First(&paintingSchool).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return err
		}
		if err := tx.Model(artist).Association("PaintingSchools").Append(paintingSchool).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, artMovement := range artist.ArtMovements {
		err := tx.Where(&model.ArtMovement{Name: artMovement.Name}).First(&artMovement).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return err
		}
		if err := tx.Model(artist).Association("ArtMovements").Append(artMovement).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// ListArtists returns a list of all the artist
func (as *ArtistStore) ListArtists() ([]*model.Artist, error) {
	var artists []*model.Artist

	err := as.db.Preload("ArtMovements").Preload("PaintingSchools").Preload("Nationalities").Preload("Paintings").Find(&artists).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return artists, err
}

// GetArtistImages returns the artist and all the images which was painted by him
func (as *ArtistStore) GetArtistImages(artistName string) (*model.Artist, error) {
	var artist model.Artist

	err := as.db.Where(&model.Artist{Name: artistName}).Preload("ArtMovements").Preload("PaintingSchools").Preload("Nationalities").Preload("Paintings").Preload("Paintings.Images").Find(&artist).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &artist, err
}

func (as *ArtistStore) GetArtists(artist *model.Artist) ([]*model.Artist, error) {
	var artists []*model.Artist

	var tx *gorm.DB = as.db

	var nationality model.Nationality
	var artistsByNationality []*model.Artist

	var artMovement model.ArtMovement
	var artistsByArtMovement []*model.Artist

	var paintingSchool model.PaintingSchool
	var artistsByPaintingSchool []*model.Artist

	var artistByName []*model.Artist

	if len(artist.Nationalities) == 1 {
		err := as.db.Where(&model.Nationality{Demonym: artist.Nationalities[0].Demonym}).First(&nationality).Error

		if err != nil {
			return nil, err
		}

		err = as.db.Select("id").Model(&nationality).Related(&artistsByNationality, "Artists").Error

		ids := make([]uint, len(artistsByNationality))
		for i, artist := range artistsByNationality {
			ids[i] = artist.ID
		}

		tx = tx.Where(ids)
	}

	if len(artist.ArtMovements) == 1 {
		err := as.db.Where(&model.ArtMovement{Name: artist.ArtMovements[0].Name}).First(&artMovement).Error

		if err != nil {
			return nil, err
		}

		err = as.db.Select("id").Model(&artMovement).Related(&artistsByArtMovement, "Artists").Error

		ids := make([]uint, len(artistsByArtMovement))
		for i, artist := range artistsByArtMovement {
			ids[i] = artist.ID
		}
		tx = tx.Where(ids)
	}

	if len(artist.PaintingSchools) == 1 {
		err := as.db.Where(&model.PaintingSchool{Name: artist.PaintingSchools[0].Name}).First(&paintingSchool).Error

		if err != nil {
			return nil, err
		}

		err = as.db.Select("id").Model(&paintingSchool).Related(&artistsByPaintingSchool, "Artists").Error

		ids := make([]uint, len(artistsByPaintingSchool))
		for i, artist := range artistsByPaintingSchool {
			ids[i] = artist.ID
		}

		tx = tx.Where(ids)
	}

	if len(artist.Name) >= 3 {
		err := as.db.Where("name LIKE ?", fmt.Sprintf("%s%s%s", "%", artist.Name, "%")).Find(&artistByName).Error

		if err != nil {
			return nil, err
		}

		ids := make([]uint, len(artistByName))
		for i, artist := range artistByName {
			ids[i] = artist.ID
		}
		tx = tx.Where(ids)
	}

	err := tx.Preload("Nationalities").Preload("PaintingSchools").Preload("ArtMovements").Preload("Paintings").Find(&artists).Error

	if err != nil {
		return nil, err
	}

	return artists, err
}
