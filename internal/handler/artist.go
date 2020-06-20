package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirtalin/democrart/internal/handler/response"
	"github.com/sirtalin/democrart/internal/model"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetArtists(c echo.Context) error {
	params := c.QueryParams()
	if len(params) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "filterRequired")
	}

	var nationalityDemonym = params.Get("nationality")
	var name string = params.Get("name")
	var paintingSchoolName string = params.Get("paintingschool")
	var artMovementName string = params.Get("artmovement")

	var artistFilter *model.Artist = new(model.Artist)

	if nationalityDemonym != "" {
		var nationality *model.Nationality = new(model.Nationality)
		nationality.Demonym = nationalityDemonym
		artistFilter.Nationalities = append(artistFilter.Nationalities, nationality)
	}

	if name != "" {
		artistFilter.Name = name
	}

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
	logrus.Debug(artistFilter)
	artists, err := h.ArtistStore.GetArtists(artistFilter)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if len(artists) == 0 {
		return c.JSON(http.StatusNoContent, artists)
	}
	return c.JSON(http.StatusOK, response.NewArtistListResponse(artists))
}
