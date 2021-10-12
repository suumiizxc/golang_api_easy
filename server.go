package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suumiizxc/golang_api/config"
	"github.com/suumiizxc/golang_api/controller"
	"github.com/suumiizxc/golang_api/middleware"
	"github.com/suumiizxc/golang_api/repository"
	"github.com/suumiizxc/golang_api/service"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	productRepository repository.ProductRepository = repository.NewProductRepository(db)
	pharmRepository   repository.PharmRepository   = repository.NewPharmRepository(db)

	jwtService     service.JWTService     = service.NewJWTService()
	userService    service.UserService    = service.NewUserService(userRepository)
	authService    service.AuthService    = service.NewAuthService(userRepository)
	productService service.ProductService = service.NewProductService(productRepository)
	pharmService   service.PharmService   = service.NewPharmService(pharmRepository)

	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	userController    controller.UserController    = controller.NewUserController(userService, jwtService)
	productController controller.ProductController = controller.NewBookController(productService, jwtService)
	pharmController   controller.PharmController   = controller.NewPharmController(pharmService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/update", userController.Update)
	}
	bookRoutes := r.Group("api/products", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", productController.All)
		bookRoutes.POST("/", productController.Insert)
		bookRoutes.GET("/:id", productController.FindByID)
		bookRoutes.PUT("/", productController.Update)
		bookRoutes.DELETE("/:id", productController.Delete)
	}
	pharmRoutes := r.Group("api/pharm")
	{
		pharmRoutes.POST("/", pharmController.Insert)
		pharmRoutes.PUT("/", pharmController.Update)
	}

	r.Run()
}
