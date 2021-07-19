package api

import (
	"github.com/gin-gonic/gin"
)

type Service struct {
	Router *gin.Engine
	Urls   map[string]*string
}

func NewService() *Service {
	return &Service{
		Router: gin.New(),
		Urls:   make(map[string]*string, 0),
	}
}
