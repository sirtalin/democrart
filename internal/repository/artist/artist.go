package artist

import "github.com/sirtalin/democrart/internal/model"

type Store interface {
	CreateArtist(*model.Artist) error
	GetArtist(string) (*model.Artist, error)
	GetArtistImages(string) (*model.Artist, error)
	ListArtists() ([]*model.Artist, error)
}
