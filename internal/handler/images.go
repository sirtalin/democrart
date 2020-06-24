package handler

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/sirtalin/democrart/internal/model"
	"github.com/sirtalin/democrart/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/verybluebot/tarinator-go"
)

func (h *Handler) GetArtistsImages(c echo.Context) error {
	// var pathTarFiles string = "/tmp"
	var locations []string

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

	artists, err := h.ArtistStore.GetArtistImages(artistFilter)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No images for that criteria"))
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	for _, artist := range artists {
		tmpDir, err := ioutil.TempDir(os.TempDir(), strings.ReplaceAll(artist.Name, " ", "_"))
		if err != nil {
			logrus.Error("Cannot create temporary directory", err)
		}

		defer os.RemoveAll(tmpDir)

		for _, painting := range artist.Paintings {

			if err != nil {
				logrus.Error("Cannot create temporary file", err)
			}

			for _, image := range painting.Images {
				imPath := path.Join(tmpDir, fmt.Sprintf("%s_%d.jpg", strings.ReplaceAll(painting.Name, " ", "_"), rand.Intn(10000)))
				err := utils.CopyFile(image.Location, imPath)
				if err == nil {
					logrus.Debugf("Copied %s in %s", image.Location, imPath)

				} else {
					logrus.Error(err)
				}
			}
		}
		locations = append(locations, tmpDir)
	}
	var file string = path.Join(fmt.Sprintf("%s%d.tar", "democrart", rand.Intn(10000)))
	err = tarinator.Tarinate(locations, file)
	if err != nil {
		logrus.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	defer os.Remove(file)

	return c.File(file)
}

func (h *Handler) GetPaintingsImages(c echo.Context) error {
	var locations []string

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

	paintings, err := h.PaintingStore.GetPaintingsImages(paintingFilter)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("No images for that criteria"))
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	for _, artist := range paintings {
		if len(artist.Paintings) != 0 {
			tmpDir, err := ioutil.TempDir(os.TempDir(), strings.ReplaceAll(artist.Name, " ", "_"))
			if err != nil {
				logrus.Error("Cannot create temporary directory", err)
			}

			defer os.RemoveAll(tmpDir)

			for _, painting := range artist.Paintings {

				if err != nil {
					logrus.Error("Cannot create temporary file", err)
				}

				for _, image := range painting.Images {
					imPath := path.Join(tmpDir, fmt.Sprintf("%s_%d.jpg", strings.ReplaceAll(painting.Name, " ", "_"), rand.Intn(10000)))
					err := utils.CopyFile(image.Location, imPath)
					if err == nil {
						logrus.Debugf("Copied %s in %s", image.Location, imPath)

					} else {
						logrus.Error(err)
					}
				}
			}
			locations = append(locations, tmpDir)
		}
	}
	var file string = path.Join(fmt.Sprintf("%s%d.tar", "democrart", rand.Intn(10000)))
	err = tarinator.Tarinate(locations, file)
	if err != nil {
		logrus.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	defer os.Remove(file)

	return c.File(file)
}
