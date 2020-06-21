package handler

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func (h *Handler) GetPaintingSchools(c echo.Context) error {
	paintingSchools, err := h.ArtistStore.GetPaintingSchools()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No painting schools found"))
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, paintingSchools)
}

func (h *Handler) GetArtistsByPaintingSchool(c echo.Context) error {
	var paintingSchool string = c.QueryParam("school")
	if paintingSchool == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Filter required")
	}
	artists, err := h.ArtistStore.GetArtistsByPaintingSchool(paintingSchool)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No records found for %s painting school", paintingSchool))
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, artists)
}
