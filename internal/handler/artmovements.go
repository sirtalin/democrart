package handler

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func (h *Handler) GetArtMovements(c echo.Context) error {
	artMovements, err := h.ArtistStore.GetArtMovements()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No art movements found"))
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, artMovements)
}

func (h *Handler) GetArtistsByArtMovement(c echo.Context) error {
	var artMovement string = c.QueryParam("movement")
	if artMovement == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Filter required")
	}
	artists, err := h.ArtistStore.GetArtistsByArtMovement(artMovement)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No records found for %s art movement", artMovement))
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, artists)
}
