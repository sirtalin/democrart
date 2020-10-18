package handler

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/sirtalin/democrart/internal/handler/response"
	"github.com/sirtalin/democrart/internal/model"
)

func (h *Handler) GetPaintings(c echo.Context) error {
	params := c.QueryParams()

	var genreName = params.Get("genre")
	var name string = params.Get("name")
	var styleName string = params.Get("style")
	var mediaName string = params.Get("media")

	if len(params) == 0 || (genreName == "" && name == "" && styleName == "" && mediaName == "") {
		return echo.NewHTTPError(http.StatusBadRequest, "Filter required")
	}

	if len(name) > 0 && len(name) < 3 {
		return echo.NewHTTPError(http.StatusBadRequest, "Name should contain 3 characters or more")
	}

	var paintingFilter *model.Painting = new(model.Painting)

	if genreName != "" {
		var genre *model.Genre = new(model.Genre)
		genre.Name = genreName
		paintingFilter.Genres = append(paintingFilter.Genres, genre)
	}

	paintingFilter.Name = name

	if styleName != "" {
		var style *model.ArtMovement = new(model.ArtMovement)
		style.Name = styleName
		paintingFilter.Styles = append(paintingFilter.Styles, style)
	}

	if mediaName != "" {
		var media *model.Media = new(model.Media)
		media.Name = mediaName
		paintingFilter.Medias = append(paintingFilter.Medias, media)
	}

	paintings, err := h.PaintingStore.GetPaintings(paintingFilter)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No records found for that criteria"))
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	if len(paintings) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No records found for that criteria"))
	}
	return c.JSON(http.StatusOK, response.NewArtistPaintingsResponse(paintings))
}
