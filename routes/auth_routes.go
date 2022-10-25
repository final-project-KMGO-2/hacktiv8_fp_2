package routes

import (
	"hacktiv8_fp_2/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, authController controller.AuthController) {
	orderRoutes := router.Group("/auth")
	{
		orderRoutes.POST("/register", authController.Register)
		orderRoutes.POST("/login", authController.Login)
	}
}
