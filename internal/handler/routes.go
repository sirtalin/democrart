package handler

import "github.com/labstack/echo"

func (h *Handler) Register(group *echo.Group) {
	artist := group.Group("/artist")
	artist.GET("/list", h.GetArtists)

	nationalities := group.Group("/nationality")
	nationalities.GET("/list", h.GetNationalities)
	nationalities.GET("/list/:nationality", h.GetArtistsByNationality)

	artMovements := group.Group("/artmovement")
	artMovements.GET("/list", h.GetArtMovements)
	artMovements.GET("/list/:artmovement", h.GetArtistsByArtMovement)

	paintingSchools := group.Group("/paintingschool")
	paintingSchools.GET("/list", h.GetPaintingSchools)
	paintingSchools.GET("/list/:paintingschool", h.GetArtistsByPaintingSchool)
}
