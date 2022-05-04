package config

import (
	"github.com/gin-gonic/gin"
)

var App *Services

type Services struct {
	R    *gin.Engine
	C    *GinConfig
	MODE string
}

func NewServices(R *gin.Engine, C *GinConfig, MODE string) *Services {
	return &Services{
		R:    R,
		C:    C,
		MODE: MODE,
	}
}

func (s *Services) GetServerMode() string {
	return s.MODE
}
