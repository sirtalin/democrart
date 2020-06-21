package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetPaintingSchools(c echo.Context) error {
	paintingSchools, err := h.ArtistStore.GetPaintingSchools()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if len(paintingSchools) == 0 {
		return c.JSON(http.StatusNoContent, paintingSchools)
	}
	return c.JSON(http.StatusOK, paintingSchools)
}

func (h *Handler) GetArtistsByPaintingSchool(c echo.Context) error {
	var paintingSchool string = c.Param("paintingschool")
	artists, err := h.ArtistStore.GetArtistsByPaintingSchool(paintingSchool)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if len(artists) == 0 {
		return c.JSON(http.StatusNoContent, artists)
	}
	return c.JSON(http.StatusOK, artists)
}
