package handler

import "github.com/labstack/echo"

func (h *Handler) Register(group *echo.Group) {
	artist := group.Group("/artist")
	artist.GET("/list", h.GetArtists)

	nationalities := group.Group("/nationalities")
	nationalities.GET("/list", h.GetNationalities)
	nationalities.GET("/list/:nationality", h.GetArtistsByNationalities)
}
