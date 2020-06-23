package painting

import "github.com/sirtalin/democrart/internal/model"

type Store interface {
	CreatePainting(*model.Artist, *model.Painting) error
	CreateImage(*model.Painting, *model.Image) error
	GetPainting(string, string) (*model.Painting, error)
	GetPaintings(*model.Painting) ([]*model.Artist, error)
}
