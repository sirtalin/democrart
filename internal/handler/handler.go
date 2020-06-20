package handler

import (
	"github.com/sirtalin/democrart/internal/repository/artist"
	"github.com/sirtalin/democrart/internal/repository/painting"
)

type Handler struct {
	ArtistStore   artist.Store
	PaintingStore painting.Store
}

func NewHandler(artistStore artist.Store, paintingStore painting.Store) *Handler {
	return &Handler{
		ArtistStore:   artistStore,
		PaintingStore: paintingStore,
	}
}
