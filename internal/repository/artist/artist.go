package artist

import "github.com/sirtalin/democrart/internal/model"

type Store interface {
	CreateArtist(*model.Artist) error
	GetArtist(string) (*model.Artist, error)
	GetArtistImages(*model.Artist) ([]*model.Artist, error)
	ListArtists() ([]*model.Artist, error)
	GetArtists(*model.Artist) ([]*model.Artist, error)
	GetNationalities() (map[string]int, error)
	GetArtMovements() (map[string]int, error)
	GetPaintingSchools() (map[string]int, error)
	GetArtistsByNationality(string) (map[string][]string, error)
	GetArtistsByArtMovement(string) (map[string][]string, error)
	GetArtistsByPaintingSchool(string) (map[string][]string, error)
}
