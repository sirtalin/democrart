package handler

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetNationalities(c echo.Context) error {
	nationalities, err := h.ArtistStore.GetNationalities()
	logrus.Debug(nationalities)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No nationalities found"))
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, nationalities)
}

func (h *Handler) GetArtistsByNationality(c echo.Context) error {
	var nationality string = c.QueryParam("nationality")
	if nationality == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Filter required")
	}
	artists, err := h.ArtistStore.GetArtistsByNationality(nationality)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No records found for %s nationality", nationality))
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, artists)
}
