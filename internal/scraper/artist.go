package scraper

import (
	"fmt"
	"strings"
	"time"

	"github.com/gertd/go-pluralize"
	"github.com/gocolly/colly"
	"github.com/sirtalin/democrart/internal/model"
	"github.com/sirtalin/democrart/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// GetArtistURLs returns a list with the urls of the authors in that art movement
func GetArtistURLs(artMovementURL string) []string {
	start := time.Now()
	var artistsURLs []string
	var wikiartProtocol string = viper.GetString("wikiart_protocol")
	var wikiartURL string = viper.GetString("wikiart_url")

	c := colly.NewCollector()

	c.OnHTML(".artist-name a", func(e *colly.HTMLElement) {
		artistURL := fmt.Sprintf("%s%s%s", wikiartProtocol, wikiartURL, e.Attr("href"))
		artistsURLs = append(artistsURLs, artistURL)
	})

	c.OnRequest(func(r *colly.Request) {
		logrus.Tracef("Visiting %s", r.URL.String())
	})

	c.Visit(artMovementURL)

	elapsed := time.Since(start)
	logrus.Infof("Get %d artist URLs from %s took %s", len(artistsURLs), artMovementURL, elapsed)

	return artistsURLs
}

// GetArtist scrap the artist information
func GetArtist(artistURL string) *model.Artist {
	var artist *model.Artist = new(model.Artist)
	var nationalities []*model.Nationality
	var artMovements []*model.ArtMovement
	var paintingSchools []*model.PaintingSchool
	var layout string = viper.GetString("time_layout_us")
	pluralize := pluralize.NewClient()

	c := colly.NewCollector()

	c.OnHTML(".wiki-layout-artist-info", func(e *colly.HTMLElement) {
		var err error

		artist.Name = utils.PrepareString(e.ChildText("h1"))
		artist.OriginalName = e.ChildText("h2[itemprop=additionalName]")
		if artist.OriginalName == "" {
			artist.OriginalName = e.ChildText("h1")
		}
		artist.BirthDate, err = time.Parse(layout, e.ChildText("span[itemprop=birthDate]"))
		if err != nil {
			logrus.Warningf("Error while parsing birth date. %s", err)
		}
		artist.DeathDate, err = time.Parse(layout, e.ChildText("span[itemprop=deathDate]"))
		if err != nil {
			logrus.Warningf("Error while parsing death date. %s", err)
		}

		e.ForEach("span[itemprop=nationality]", func(_ int, el *colly.HTMLElement) {
			var nationality *model.Nationality = new(model.Nationality)
			nationality.Demonym = pluralize.Singular(strings.ToLower(el.Text))
			nationalities = append(nationalities, nationality)
		})
		artist.Nationalities = nationalities
		e.ForEach("li.dictionary-values", func(_ int, el *colly.HTMLElement) {
			var text string = utils.TrimAllSpaces(el.Text)
			var textArray []string = strings.Split(text, ": ")
			if strings.Contains(text, "Art Movement") {
				for _, artMovementName := range strings.Split(textArray[1], ", ") {
					var artMovement *model.ArtMovement = new(model.ArtMovement)
					artMovement.Name = utils.PrepareString(artMovementName)
					artMovements = append(artMovements, artMovement)
				}
				artist.ArtMovements = artMovements
			}
			if strings.Contains(text, "Painting School") {
				for _, paintingSchoolName := range strings.Split(textArray[1], ", ") {
					var paintingSchool *model.PaintingSchool = new(model.PaintingSchool)
					paintingSchool.Name = utils.PrepareString(paintingSchoolName)
					paintingSchools = append(paintingSchools, paintingSchool)
				}
				artist.PaintingSchools = paintingSchools
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		logrus.Tracef("Visiting %s", r.URL.String())
	})

	c.Visit(artistURL)

	return artist
}
