package main

import (
	"hacktiv8_fp_2/config"
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/repository"
	"hacktiv8_fp_2/routes"
	"hacktiv8_fp_2/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	var (
		db             *gorm.DB                  = config.SetupDatabaseConnection()
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		socmedRepository repository.SocmedRepository = repository.NewSocmedRepository(db)
		// productRepository repository.ProductRepository = repository.NewProductRepository(db)

		jwtService  service.JWTService  = service.NewJWTService()
		userService service.UserService = service.NewUserService(userRepository)
		authService service.AuthService = service.NewAuthService(userRepository)
		socmedService service.SocmedService = service.NewSocmedService(socmedRepository)

		// productController controller.ProductController = controller.NewProductController(productService, jwtService)
		authController controller.AuthController = controller.NewAuthController(userService, authService, jwtService)
		socmedController controller.SocmedController = controller.NewSocmedController(userService, socmedService, jwtService)
	)

	defer config.CloseDatabaseConnection(db)

	server := gin.Default()

	routes.AuthRoutes(server, authController);
	routes.SocMedRoutes(server, socmedController);
	// routes.ProductRoutes(server, productController, jwtService, productService)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
