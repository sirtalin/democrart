package handler

import "github.com/labstack/echo"

func (h *Handler) Register(group *echo.Group) {
	artist := group.Group("/artists")
	artist.GET("/list", h.GetArtists)

	nationalities := artist.Group("/nationalities")
	nationalities.GET("", h.GetNationalities)
	nationalities.GET("/list", h.GetArtistsByNationality)

	artMovements := artist.Group("/movements")
	artMovements.GET("", h.GetArtMovements)
	artMovements.GET("/list", h.GetArtistsByArtMovement)

	paintingSchools := artist.Group("/schools")
	paintingSchools.GET("", h.GetPaintingSchools)
	paintingSchools.GET("/list", h.GetArtistsByPaintingSchool)
}
