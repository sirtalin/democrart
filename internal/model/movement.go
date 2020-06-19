package model

import "github.com/jinzhu/gorm"

type ArtMovement struct {
	gorm.Model
	Name      string      `gorm:"type:varchar(64);not null;unique" json:"name"`
	Artists   []*Artist   `gorm:"many2many:artist_movements"`
	Paintings []*Painting `gorm:"many2many:painting_styles"`
}

type ArtistMovement struct {
	ArtistID      uint         `sql:"type:int REFERENCES artists(id)" json:"artist_id"`
	Artist        *Artist      `gorm:"foreignkey:id"`
	ArtMovementID uint         `sql:"type:int REFERENCES art_movements(id)" json:"art_movement_id"`
	ArtMovement   *ArtMovement `gorm:"foreignkey:id"`
}

type PaintingStyle struct {
	PaintingID    uint         `sql:"type:int REFERENCES paintings(id)" json:"painting_id"`
	Painting      *Painting    `gorm:"foreignkey:id"`
	ArtMovementID uint         `sql:"type:int REFERENCES art_movements(id)" json:"art_movement_id"`
	ArtMovement   *ArtMovement `gorm:"foreignkey:id"`
}
