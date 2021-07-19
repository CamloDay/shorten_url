package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shorten_url/api/handlers/context"
)

func Redirect(c *gin.Context) {
	short, ok := c.Params.Get("short_link")
	if !ok || short == "" {
		c.JSON(http.StatusBadRequest, BadRequestResponse())
		return
	}

	ctx := context.GetContext(c)
	original, err := MapOriginal(ctx.Urls, short)
	if err != nil {
		c.JSON(http.StatusNotFound, NotFoundResponse())
		return
	}

	c.Redirect(http.StatusMovedPermanently, *original)
}

func MapOriginal(urls map[string]*string, short string) (*string, error) {
	original := urls[short]
	if original == nil {
		return nil, fmt.Errorf(StatusDescriptions[StatusNotFound])
	}
	return original, nil
}
