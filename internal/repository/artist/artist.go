package artist

import "github.com/sirtalin/democrart/internal/model"

type Store interface {
	CreateArtist(*model.Artist) error
	GetArtist(string) (*model.Artist, error)
	GetArtistImages(*model.Artist) ([]*model.Artist, error)
	ListArtists() ([]*model.Artist, error)
	GetArtists(*model.Artist) ([]*model.Artist, error)
	GetNationalities() ([]model.Result, error)
	GetArtMovements() ([]model.Result, error)
	GetPaintingSchools() ([]model.Result, error)
	GetArtistsByNationality(string) (map[string][]string, error)
	GetArtistsByArtMovement(string) (map[string][]string, error)
	GetArtistsByPaintingSchool(string) (map[string][]string, error)
}
