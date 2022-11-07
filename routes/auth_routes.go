package routes

import (
	"hacktiv8_fp_2/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, authController controller.AuthController) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
	}
}
