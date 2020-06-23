package response

import (
	"time"

	"github.com/sirtalin/democrart/internal/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ArtistResponse struct {
	Name            string   `json:"name"`
	OriginalName    string   `json:"original_name"`
	Nationalities   []string `json:"nationalities"`
	PaintingSchools []string `json:"painting_schools"`
	ArtMovements    []string `json:"art_movements"`
	Paintings       []string `json:"paintings"`
	BirthDate       string   `json:"birth_date"`
	DeathDate       string   `json:"death_date"`
}

type PaintingResponse struct {
	Name         string   `json:"name"`
	OriginalName string   `json:"original_name"`
	Width        int      `json:"width"`
	Height       int      `json:"height"`
	Genres       []string `json:"genres"`
	Styles       []string `json:"styles"`
	Medias       []string `json:"medias"`
}

func NewArtistListResponse(artists []*model.Artist) []*ArtistResponse {
	var artistResponse *ArtistResponse
	var artistResponseList []*ArtistResponse
	for _, artist := range artists {
		artistResponse = NewArtistResponse(artist)
		artistResponseList = append(artistResponseList, artistResponse)
	}
	return artistResponseList
}

func formatDate(date time.Time) string {
	var layout string = viper.GetString("time_layout_us")
	if date.IsZero() {
		return "Unknown"
	}
	return date.Format(layout)
}

func NewArtistResponse(artist *model.Artist) *ArtistResponse {
	var artistResponse *ArtistResponse = new(ArtistResponse)
	artistResponse.Name = artist.Name
	artistResponse.OriginalName = artist.OriginalName
	for _, nationality := range artist.Nationalities {
		artistResponse.Nationalities = append(artistResponse.Nationalities, nationality.Demonym)
	}
	for _, paintingSchool := range artist.PaintingSchools {
		artistResponse.PaintingSchools = append(artistResponse.PaintingSchools, paintingSchool.Name)
	}
	for _, artMovement := range artist.ArtMovements {
		artistResponse.ArtMovements = append(artistResponse.ArtMovements, artMovement.Name)
	}
	for _, painting := range artist.Paintings {
		artistResponse.Paintings = append(artistResponse.Paintings, painting.Name)
	}

	artistResponse.BirthDate = formatDate(artist.BirthDate)
	artistResponse.DeathDate = formatDate(artist.DeathDate)

	return artistResponse
}

func NewPaintingMapResponse(artists []*model.Artist) map[string][]*PaintingResponse {
	var paintingMapResponse map[string][]*PaintingResponse = make(map[string][]*PaintingResponse)
	var count uint
	for _, artist := range artists {
		if len(artist.Paintings) > 0 {
			for _, painting := range artist.Paintings {
				paintingMapResponse[artist.Name] = append(paintingMapResponse[artist.Name], NewPaintingResponse(painting))
			}
		} else {
			count++
		}
	}
	logrus.Info(count, " artists without paintings")

	return paintingMapResponse
}

func NewPaintingResponse(painting *model.Painting) *PaintingResponse {
	var paintingResponse *PaintingResponse = new(PaintingResponse)
	paintingResponse.Name = painting.Name
	paintingResponse.OriginalName = painting.OriginalName
	paintingResponse.Width = painting.Width
	paintingResponse.Height = painting.Height
	for _, genre := range painting.Genres {
		paintingResponse.Genres = append(paintingResponse.Genres, genre.Name)
	}
	for _, style := range painting.Styles {
		paintingResponse.Styles = append(paintingResponse.Styles, style.Name)
	}
	for _, media := range painting.Medias {
		paintingResponse.Medias = append(paintingResponse.Medias, media.Name)
	}

	return paintingResponse
}
