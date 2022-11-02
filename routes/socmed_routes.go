package routes

import (
	"hacktiv8_fp_2/controller"

	"github.com/gin-gonic/gin"
)

func SocMedRoutes(router *gin.Engine, socmedController controller.SocmedController) {
	smRoutes := router.Group("/socialmedias")
	{
		smRoutes.POST("/", socmedController.PostSocmed)
		smRoutes.GET("/", socmedController.GetSocmed)
		smRoutes.PUT("/:id", socmedController.UpdateSocmedById)
		smRoutes.DELETE("/:id", socmedController.DeleteSocmedById)
	}
}
