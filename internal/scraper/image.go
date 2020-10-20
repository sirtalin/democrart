package scraper

import (
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/sirtalin/democrart/internal/model"
	"github.com/sirtalin/democrart/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func downloadImage(paintingImageURI string, artistName string) string {
	var absImagePath string
	var err error

	c := colly.NewCollector(
		colly.AllowedDomains("uploads1.wikiart.org", "uploads2.wikiart.org",
			"uploads3.wikiart.org", "uploads4.wikiart.org",
			"uploads5.wikiart.org", "uploads6.wikiart.org",
			"uploads7.wikiart.org", "uploads8.wikiart.org",
			"uploads9.wikiart.org", "uploads0.wikiart.org"),
	)

	c.OnResponse(func(r *colly.Response) {
		if strings.Index(r.Headers.Get("Content-Type"), "image") > -1 {
			imagePath := path.Join(viper.GetString("paintings_directory"), artistName)
			utils.CreateDirectory(imagePath)
			imagePath = path.Join(imagePath, strings.ReplaceAll(r.FileName(), "-", "_"))
			if !utils.FileExists(imagePath) {
				r.Save(imagePath)
				logrus.Tracef("Saving image from uri %s in %s", paintingImageURI, imagePath)
				absImagePath, err = filepath.Abs(imagePath)
				if err != nil {
					logrus.Errorf("Error obtaining absolute path for %s. %s", imagePath, err)
				}
			}

		}
	})

	c.Visit(paintingImageURI)

	return absImagePath
}

// GetImage returns the location and size of the downloaded painting image
func GetImage(paintingURL string) *model.Image {
	var image *model.Image = new(model.Image)
	var artistName string

	c := colly.NewCollector(
		colly.AllowedDomains("www.wikiart.org"),
	)

	c.OnHTML("span.max-resolution", func(e *colly.HTMLElement) {
		maxResolution := strings.Split(strings.TrimSuffix(e.Text, "px"), "x")
		image.Width, _ = strconv.Atoi(maxResolution[0])
		image.Height, _ = strconv.Atoi(maxResolution[1])
	})

	c.OnHTML("h5[itemprop=creator] a", func(e *colly.HTMLElement) {
		artistName = utils.ToSnakeCase(e.Text)
	})

	c.OnHTML("img.ms-zoom-cursor", func(e *colly.HTMLElement) {
		image.Location = downloadImage(strings.Split(e.Attr("src"), "!")[0], artistName)
	})

	c.OnRequest(func(r *colly.Request) {
		logrus.Tracef("Visiting %s", r.URL.String())
	})

	c.Visit(paintingURL)

	return image
}
