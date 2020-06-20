package response

import (
	"time"

	"github.com/sirtalin/democrart/internal/model"
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