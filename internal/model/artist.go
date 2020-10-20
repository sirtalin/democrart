package model

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Artist struct {
	gorm.Model
	Name            string            `gorm:"type:varchar(128);not null;unique" json:"name"`
	OriginalName    string            `gorm:"type:varchar(128)" json:"original_name"`
	Nationalities   []*Nationality    `gorm:"many2many:artists_nationalities;association_autocreate:false"`
	PaintingSchools []*PaintingSchool `gorm:"many2many:artists_schools;association_autocreate:false"`
	ArtMovements    []*ArtMovement    `gorm:"many2many:artists_movements;association_autocreate:false"`
	Paintings       []*Painting       `gorm:"foreignkey:artist_id"`
	BirthDate       time.Time         `gorm:"type:date" json:"birth_date"`
	DeathDate       time.Time         `gorm:"type:date" json:"death_date"`
}

type PaintingSchool struct {
	gorm.Model
	Name    string    `gorm:"type:varchar(64);not null;unique" json:"name"`
	Artists []*Artist `gorm:"many2many:artists_schools"`
}

type Nationality struct {
	gorm.Model
	Demonym string    `gorm:"type:varchar(32);not null;unique" json:"demonym"`
	Artists []*Artist `gorm:"many2many:artists_nationalities"`
}

type ArtistsSchool struct {
	ArtistID         uint            `sql:"type:int REFERENCES artists(id)" json:"artist_id"`
	Artist           *Artist         `gorm:"foreignkey:id"`
	PaintingSchoolID uint            `sql:"type:int REFERENCES painting_schools(id)" json:"painting_school_id"`
	PaintingSchool   *PaintingSchool `gorm:"foreignkey:id"`
}

type ArtistsNationality struct {
	ArtistID      uint         `sql:"type:int REFERENCES artists(id)" json:"artist_id"`
	Artist        *Artist      `gorm:"foreignkey:id"`
	NationalityID uint         `sql:"type:int REFERENCES nationality(id)" json:"nationality_id"`
	Nationality   *Nationality `gorm:"foreignkey:id"`
}

type ArtistCSV struct {
	ID              int             `csv:"id"`
	Name            string          `csv:"name"`
	OriginalName    string          `csv:"original_name"`
	Nationalities   Nationalities   `csv:"nationalities"`
	PaintingSchools PaintingSchools `csv:"painting_schools"`
	ArtMovements    ArtMovements    `csv:"art_movements"`
	BirthDate       Date            `csv:"birth_date"`
	DeathDate       Date            `csv:"death_date"`
}

type Nationalities struct {
	Nationalities []string
}

func (nationalities *Nationalities) UnmarshalCSV(csv string) (err error) {
	for _, nationality := range strings.Split(csv, ";") {
		nationalities.Nationalities = append(nationalities.Nationalities, nationality)
	}
	return nil
}

type PaintingSchools struct {
	PaintingSchools []string
}

func (paintingSchools *PaintingSchools) UnmarshalCSV(csv string) (err error) {
	for _, paintingSchool := range strings.Split(csv, ";") {
		paintingSchools.PaintingSchools = append(paintingSchools.PaintingSchools, paintingSchool)
	}
	return nil

}

type ArtMovements struct {
	ArtMovements []string
}

func (paintingSchools *ArtMovements) UnmarshalCSV(csv string) (err error) {
	for _, artMovement := range strings.Split(csv, ";") {
		paintingSchools.ArtMovements = append(paintingSchools.ArtMovements, artMovement)
	}
	return nil

}
