package routes

import (
	"hacktiv8_fp_2/controller"

	"github.com/gin-gonic/gin"
)

func SocMedRoutes(router *gin.Engine) {
	smRoutes := router.Group("/socialmedias")
	{
		smRoutes.POST("/", )
		smRoutes.GET("/",)
		smRoutes.PUT("/{id}")
		smRoutes.DELETE("/{id}")
	}
}
