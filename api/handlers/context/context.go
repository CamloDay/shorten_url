package context

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	Urls map[string]*string
	Addr string
}

func SetContext(c *gin.Context, a *Context) {
	c.Set("apiContext", a)
}

func GetContext(c *gin.Context) *Context {
	apiEnv := c.MustGet("apiContext").(*Context)
	return apiEnv
}
