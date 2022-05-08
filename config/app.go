package config

import (
	"github.com/gin-gonic/gin"
)

var App *Services

type Services struct {
	R            *gin.Engine
	C            *GinConfig
	MODE         string
	AUTH_SERVICE string
}

func NewServices(R *gin.Engine, C *GinConfig, MODE string, authService string) *Services {
	return &Services{
		R:            R,
		C:            C,
		MODE:         MODE,
		AUTH_SERVICE: authService,
	}
}

func (s *Services) GetServerMode() string {
	return s.MODE
}
