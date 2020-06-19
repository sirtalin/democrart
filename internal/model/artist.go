package model

import (
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
