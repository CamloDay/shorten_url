package api

import (
	"github.com/gin-gonic/gin"
	"shorten_url/api/handlers/context"
)

func AddContext(s *Service, addr string) gin.HandlerFunc {
	return func(c *gin.Context) {
		context.SetContext(c, &context.Context{
			Urls: s.Urls,
			Addr: addr,
		})
		c.Next()
	}
}
