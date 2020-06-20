package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetArtMovements(c echo.Context) error {
	artMovements, err := h.ArtistStore.GetArtMovements()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if len(artMovements) == 0 {
		return c.JSON(http.StatusNoContent, artMovements)
	}
	return c.JSON(http.StatusOK, artMovements)
}

func (h *Handler) GetArtistsByArtMovement(c echo.Context) error {
	var artMovement string = c.Param("artmovement")
	artists, err := h.ArtistStore.GetArtistsByArtMovement(artMovement)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if len(artists) == 0 {
		return c.JSON(http.StatusNoContent, artists)
	}
	return c.JSON(http.StatusOK, artists)
}
