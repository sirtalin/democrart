package handler

import "github.com/labstack/echo"

func (h *Handler) Register(group *echo.Group) {
	artists := group.Group("/artists")
	artists.GET("/list", h.GetArtists)
	artists.GET("/images", h.GetArtistsImages)

	nationalities := artists.Group("/nationalities")
	nationalities.GET("", h.GetNationalities)
	nationalities.GET("/list", h.GetArtistsByNationality)

	artMovements := artists.Group("/movements")
	artMovements.GET("", h.GetArtMovements)
	artMovements.GET("/list", h.GetArtistsByArtMovement)

	paintingSchools := artists.Group("/schools")
	paintingSchools.GET("", h.GetPaintingSchools)
	paintingSchools.GET("/list", h.GetArtistsByPaintingSchool)

	paintings := group.Group("/paintings")
	paintings.GET("/list", h.GetPaintings)
}
