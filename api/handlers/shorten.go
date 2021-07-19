package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"shorten_url/api/handlers/context"
)

type ShortRequest struct {
	Original string `json:"original" binding:"required"`
}

type ShortResponse struct {
	StandardResponse
	Short string `json:"short"`
}

func ShortenUrl(c *gin.Context) {
	var req ShortRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, BadRequestResponse())
		return
	}

	ctx := context.GetContext(c)
	short := Shorten(&req.Original, ctx.Urls, ctx.Addr)

	c.JSON(http.StatusOK, &ShortResponse{
		StandardResponse: StandardResponse{
			StatusCode:  StatusOk,
			Description: StatusDescriptions[StatusOk],
		},
		Short: short,
	})
}

func Shorten(original *string, urls map[string]*string, addr string) string {
	u := uuid.NewV4().String()
	short := "http://" + addr + "/" + u
	urls[u] = original
	return short
}
