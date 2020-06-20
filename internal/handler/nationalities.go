package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetNationalities(c echo.Context) error {
	nationalities, err := h.ArtistStore.GetNationalities()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if len(nationalities) == 0 {
		return c.JSON(http.StatusNoContent, nationalities)
	}
	return c.JSON(http.StatusOK, nationalities)
}

func (h *Handler) GetArtistsByNationality(c echo.Context) error {
	var nationality string = c.Param("nationality")
	artists, err := h.ArtistStore.GetArtistsByNationality(nationality)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if len(artists) == 0 {
		return c.JSON(http.StatusNoContent, artists)
	}
	return c.JSON(http.StatusOK, artists)
}
