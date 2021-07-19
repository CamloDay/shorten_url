package api

import (
	"shorten_url/api/handlers"
)

func CreateRoutes(s *Service, addr string) {
	vX := s.Router.Group("/", AddContext(s, addr))

	vX.POST("/shorten", handlers.ShortenUrl)
	vX.GET("/:short_link", handlers.Redirect)
}
