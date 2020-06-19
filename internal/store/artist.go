package store

import (
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
