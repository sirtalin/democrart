package scraper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/sirtalin/democrart/internal/model"
	"github.com/sirtalin/democrart/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// GetPaintingsURLs returns a list with the URLs of the paintings of wikiart
func GetPaintingsURLs(artistURL string) []string {
	var paintingsURLs []string
	var wikiartProtocol string = viper.GetString("wikiart_protocol")
	var wikiartURL string = viper.GetString("wikiart_url")

	c := colly.NewCollector()

	c.OnHTML(".painting-list-text-row a", func(e *colly.HTMLElement) {
		var paintingURL string = fmt.Sprintf("%s%s%s", wikiartProtocol, wikiartURL, e.Attr("href"))
		paintingsURLs = append(paintingsURLs, paintingURL)
	})

	c.OnRequest(func(r *colly.Request) {
		logrus.Tracef("Visiting %s", r.URL.String())
	})

	c.Visit(artistURL)

	return paintingsURLs
}

// GetPainting obtain information of the painting URL passed as argument
func GetPainting(paintingURL string) *model.Painting {
	var painting *model.Painting = new(model.Painting)
	var styles []*model.ArtMovement
	var genres []*model.Genre
	var medias []*model.Media

	c := colly.NewCollector()

	c.OnHTML(".wiki-layout-artist-info h3", func(e *colly.HTMLElement) {
		painting.OriginalName = e.Text
		painting.Name = utils.PrepareString(e.Text)
	})

	c.OnHTML("a.copyright", func(e *colly.HTMLElement) {
		painting.Copyright = e.Text
	})

	c.OnHTML(".wiki-layout-artist-info li", func(e *colly.HTMLElement) {
		text := utils.TrimAllSpaces(e.Text)
		textArray := strings.Split(text, ": ")
		if len(textArray) > 1 {
			text = textArray[1]
		}

		if strings.Contains(e.Text, "Original Title") {
			painting.OriginalName = text
		}

		if strings.Contains(e.Text, "Style") {
			for _, styleName := range strings.Split(text, ", ") {
				var style *model.ArtMovement = new(model.ArtMovement)
				style.Name = utils.PrepareString(styleName)
				styles = append(styles, style)
			}
			painting.Styles = styles
		}

		if strings.Contains(e.Text, "Genre") {
			for _, genreName := range strings.Split(text, ", ") {
				var genre *model.Genre = new(model.Genre)
				genre.Name = utils.PrepareString(genreName)
				genres = append(genres, genre)
			}
			painting.Genres = genres
		}

		if strings.Contains(e.Text, "Media") {
			for _, mediaName := range strings.Split(text, ", ") {
				var media *model.Media = new(model.Media)
				media.Name = utils.PrepareString(mediaName)
				medias = append(medias, media)
			}
			painting.Medias = medias
		}

		if strings.Contains(e.Text, "Dimensions") {
			dimension := strings.Split(strings.TrimSuffix(text, " cm"), " x ")
			painting.Width, _ = strconv.Atoi(dimension[0])
			painting.Height, _ = strconv.Atoi(dimension[1])
		}
	})

	c.OnRequest(func(r *colly.Request) {
		logrus.Tracef("Visiting %s", r.URL.String())
	})

	c.Visit(paintingURL)

	return painting
}
