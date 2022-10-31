package controller

import (
	"hacktiv8_fp_2/service"

	"github.com/gin-gonic/gin"
)

type SocmedController interface {
	PostSocmed(ctx *gin.Context)
	GetSocmed(ctx *gin.Context)
	UpdateSocmedById(ctx *gin.Context)
	DeleteSocmedById(ctx *gin.Context)
}

type socmedController struct {
	userService service.UserService
	jwtService  service.JWTService
	socmedService service.SocmedService
}


func NewSocmedController(us service.UserService, ss service.SocmedService, js service.JWTService) SocmedController {
	return &socmedController {
		userService: us,
		socmedService: ss,
		jwtService:  js,
	}
}


func (sc *socmedController) PostSocmed(ctx *gin.Context) {}
func (sc *socmedController) GetSocmed(ctx *gin.Context) {}
func (sc *socmedController) UpdateSocmedById(ctx *gin.Context) {}
func (sc *socmedController) DeleteSocmedById(ctx *gin.Context) {}
