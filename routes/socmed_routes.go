package routes

import (
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/middleware"
	"hacktiv8_fp_2/service"

	"github.com/gin-gonic/gin"
)

func SocMedRoutes(router *gin.Engine, socmedController controller.SocmedController, jwtService service.JWTService) {
	smRoutes := router.Group("/socialmedias")
	{
		smRoutes.POST("/", middleware.Authenticate(jwtService), socmedController.PostSocmed)
		smRoutes.GET("", socmedController.GetSocmed)
		smRoutes.PUT("/:id", socmedController.UpdateSocmedById)
		smRoutes.DELETE("/:id", socmedController.DeleteSocmedById)
	}
}
