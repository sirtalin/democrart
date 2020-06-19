package model

import (
	"github.com/jinzhu/gorm"
)

type Painting struct {
	gorm.Model
	Name         string         `gorm:"type:varchar(128);not null;unique_index:artist_painting" json:"name"`
	OriginalName string         `gorm:"type:varchar(128)" json:"original_name"`
	Artist       *Artist        `gorm:"foreignkey:id"`
	ArtistID     uint           `gorm:"unique_index:artist_painting" sql:"type:int REFERENCES artists(id)" json:"artist_id"`
	Width        int            `gorm:"type:int" json:"width"`
	Height       int            `gorm:"type:int" json:"height"`
	Copyright    string         `gorm:"type:varchar(64)" json:"copyright"`
	Genres       []*Genre       `gorm:"many2many:paintings_genres;association_autocreate:false"`
	Styles       []*ArtMovement `gorm:"many2many:paintings_styles;association_autocreate:false"`
	Medias       []*Media       `gorm:"many2many:paintings_medias;association_autocreate:false"`
	Images       []*Image       `gorm:"foreignkey:painting_id"`
}

type Genre struct {
	gorm.Model
	Name      string      `gorm:"type:varchar(64);not null;unique" json:"name"`
	Paintings []*Painting `gorm:"many2many:paintings_genres"`
}

type Media struct {
	gorm.Model
	Name      string      `gorm:"type:varchar(32);not null;unique" json:"name"`
	Paintings []*Painting `gorm:"many2many:paintings_medias"`
}

type Image struct {
	gorm.Model
	Location   string    `gorm:"type:varchar(255);not null;unique" json:"name"`
	Width      int       `gorm:"type:int;" json:"width"`
	Height     int       `gorm:"type:int;" json:"height"`
	Painting   *Painting `gorm:"foreignkey:id"`
	PaintingID uint      `sql:"type:int REFERENCES paintings(id)" json:"painting_id"`
}

type PaintingsGenre struct {
	PaintingID uint      `sql:"type:int REFERENCES paintings(id)" json:"painting_id"`
	Painting   *Painting `gorm:"foreignkey:id"`
	GenreID    uint      `sql:"type:int REFERENCES genres(id)" json:"genre_id"`
	Genre      *Genre    `gorm:"foreignkey:id"`
}

type PaintingsMedias struct {
	PaintingID uint      `sql:"type:int REFERENCES paintings(id)" json:"painting_id"`
	Painting   *Painting `gorm:"foreignkey:id"`
	MediaID    uint      `sql:"type:int REFERENCES media(id)" json:"media_id"`
	Media      *Genre    `gorm:"foreignkey:id"`
}

// Validate returns if the painting is free for use or not
func (p *Painting) Valid() bool {
	return p.Copyright == "Public domain" || p.Copyright == "Fair Use"
}
