package handler

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/sirtalin/democrart/internal/handler/response"
	"github.com/sirtalin/democrart/internal/model"
)

func (h *Handler) GetArtists(c echo.Context) error {
	params := c.QueryParams()

	var nationalityDemonym = params.Get("nationality")
	var name string = params.Get("name")
	var paintingSchoolName string = params.Get("paintingschool")
	var artMovementName string = params.Get("artmovement")

	if len(params) == 0 || (nationalityDemonym == "" && name == "" && paintingSchoolName == "" && artMovementName == "") {
		return echo.NewHTTPError(http.StatusBadRequest, "Filter required")
	}

	if len(name) > 0 && len(name) < 3 {
		return echo.NewHTTPError(http.StatusBadRequest, "Name should contain 3 characters or more")
	}

	var artistFilter *model.Artist = new(model.Artist)

	if nationalityDemonym != "" {
		var nationality *model.Nationality = new(model.Nationality)
		nationality.Demonym = nationalityDemonym
		artistFilter.Nationalities = append(artistFilter.Nationalities, nationality)
	}

	artistFilter.Name = name

	if paintingSchoolName != "" {
		var paintingSchool *model.PaintingSchool = new(model.PaintingSchool)
		paintingSchool.Name = paintingSchoolName
		artistFilter.PaintingSchools = append(artistFilter.PaintingSchools, paintingSchool)
	}

	if artMovementName != "" {
		var artMovement *model.ArtMovement = new(model.ArtMovement)
		artMovement.Name = artMovementName
		artistFilter.ArtMovements = append(artistFilter.ArtMovements, artMovement)
	}

	artists, err := h.ArtistStore.GetArtists(artistFilter)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No records found for that criteria"))
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	if len(artists) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No records found for that criteria"))
	}
	return c.JSON(http.StatusOK, response.NewArtistListResponse(artists))
}
