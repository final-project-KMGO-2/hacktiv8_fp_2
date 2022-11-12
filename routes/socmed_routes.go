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
		smRoutes.POST("", middleware.Authenticate(jwtService), socmedController.PostSocmed)
		smRoutes.GET("", middleware.Authenticate(jwtService), socmedController.GetSocmed)
		smRoutes.PUT("/:socialMediaId", middleware.Authenticate(jwtService), socmedController.UpdateSocmedById)
		smRoutes.DELETE("/:socialMediaId", middleware.Authenticate(jwtService), socmedController.DeleteSocmedById)
	}
}
