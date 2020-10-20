package scraper

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/sirtalin/democrart/internal/handler"
	"github.com/sirtalin/democrart/internal/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// GetArtMovementURLs returns a list with the URLs of the art art movements of wikiart
func GetArtMovementURLs() []string {
	start := time.Now()
	var artMovementURLs []string
	var wikiartProtocol string = viper.GetString("wikiart_protocol")
	var wikiartURL string = viper.GetString("wikiart_url")
	var scraperListBeginning int = viper.GetInt("scraper_list_beginning")

	c := colly.NewCollector()

	c.OnHTML("ul[class=dictionaries-list] li[class=dottedItem] a", func(e *colly.HTMLElement) {
		var href string = e.Attr("href")
		if href != "" {
			artMovementURL := fmt.Sprintf("%s%s%s", wikiartProtocol, wikiartURL, href)
			artMovementURLs = append(artMovementURLs, artMovementURL)
		}

	})

	c.OnRequest(func(r *colly.Request) {
		logrus.Tracef("Visiting %s", r.URL.String())
	})

	c.Visit("https://www.wikiart.org/en/artists-by-art-movement")

	elapsed := time.Since(start)
	logrus.Infof("Get %d art movements URL took %s", len(artMovementURLs), elapsed)

	return artMovementURLs[scraperListBeginning:]
}

func getArtMovementNameFromURL(artMovementURL string) string {
	var artMovementURLSplit []string = strings.Split(artMovementURL, "/")
	var artMovementName string = artMovementURLSplit[len(artMovementURLSplit)-1]
	return strings.Split(artMovementName, "#")[0]
}

// GetArtMovement scrap all the artist an paintings from an art movement and insert it on the database
func GetArtMovement(artMovementURL string, scrapartHandler *handler.Handler) {
	start := time.Now()
	var numPaintings int
	var painting *model.Painting
	var artistPaintingsURL string
	var paintingsURLs []string
	var artist *model.Artist
	var image *model.Image
	var artistsURLs []string = GetArtistURLs(artMovementURL)
	var artMovementName string = getArtMovementNameFromURL(artMovementURL)
	for _, artistURL := range artistsURLs {
		artist = GetArtist(artistURL)
		err := scrapartHandler.ArtistStore.CreateArtist(artist)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"movement": artMovementName,
				"artist":   artist.Name,
			}).Error(err)
			if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
				artist, _ = scrapartHandler.ArtistStore.GetArtist(artist.Name)
				logrus.Infof("Artist duplicated. Get artist %s with ID %d", artist.Name, artist.ID)
			}
		}

		logrus.WithFields(logrus.Fields{
			"movement": artMovementName,
		}).Info(artist)

		artistPaintingsURL = fmt.Sprintf("%s/all-works/text-list", artistURL)
		paintingsURLs = GetPaintingsURLs(artistPaintingsURL)

		for _, paintingURL := range paintingsURLs {
			painting = GetPainting(paintingURL)
			if painting.Valid() {
				image = GetImage(paintingURL)
				err := scrapartHandler.PaintingStore.CreatePainting(artist, painting)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"movement": artMovementName,
						"artist":   artist.Name,
						"painting": painting.Name,
					}).Error(err)
					if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
						painting, _ = scrapartHandler.PaintingStore.GetPainting(artist.Name, painting.Name)
						logrus.Infof("Painting duplicated. Get painting %s with ID %d", painting.Name, painting.ID)
					}
				}

				logrus.WithFields(logrus.Fields{
					"movement": artMovementName,
					"artist":   artist.Name,
				}).Info(painting)

				if image.Location != "" {
					err = scrapartHandler.PaintingStore.CreateImage(painting, image)
					if err != nil {
						logrus.WithFields(logrus.Fields{
							"movement": artMovementName,
							"artist":   artist.Name,
							"painting": painting.Name,
							"location": image.Location,
						}).Error(err)
					}

					logrus.WithFields(logrus.Fields{
						"movement": artMovementName,
						"artist":   artist.Name,
						"painting": painting.Name,
					}).Info(image)

					numPaintings++
				}
			}
		}
	}

	elapsed := time.Since(start)
	logrus.Infof("Get the %s with its %d artists and its %d paintings took %s", artMovementURL, len(artistsURLs), numPaintings, elapsed)
}
