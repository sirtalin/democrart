package handler

import "github.com/labstack/echo"

func (h *Handler) Register(group *echo.Group) {
	artist := group.Group("/artist")
	artist.GET("/list", h.GetArtists)
}
