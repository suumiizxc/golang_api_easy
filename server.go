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
	db                   *gorm.DB                        = config.SetupDatabaseConnection()
	userRepository       repository.UserRepository       = repository.NewUserRepository(db)
	productRepository    repository.ProductRepository    = repository.NewProductRepository(db)
	pharmRepository      repository.PharmRepository      = repository.NewPharmRepository(db)
	doctorRepository     repository.DoctorRepository     = repository.NewDoctorRepository(db)
	pharmacistRepository repository.PharmacistRepository = repository.NewPharmacistRepository(db)
	orderRepository      repository.OrderRepository      = repository.NewOrderRepository(db)

	jwtService            service.JWTService            = service.NewJWTService()
	userService           service.UserService           = service.NewUserService(userRepository)
	authService           service.AuthService           = service.NewAuthService(userRepository)
	productService        service.ProductService        = service.NewProductService(productRepository)
	pharmService          service.PharmService          = service.NewPharmService(pharmRepository)
	doctorService         service.DoctorService         = service.NewDoctorService(doctorRepository)
	authDoctorService     service.AuthDoctorService     = service.NewAuthDoctorService(doctorRepository)
	pharmacistService     service.PharmacistService     = service.NewPharmacistService(pharmacistRepository)
	authPharmacistService service.AuthPharmacistService = service.NewAuthPharmacistService(pharmacistRepository)
	orderService          service.OrderService          = service.NewOrderService(orderRepository)

	authController           controller.AuthController           = controller.NewAuthController(authService, jwtService)
	userController           controller.UserController           = controller.NewUserController(userService, jwtService)
	productController        controller.ProductController        = controller.NewBookController(productService, jwtService)
	pharmController          controller.PharmController          = controller.NewPharmController(pharmService)
	authDoctorController     controller.AuthDoctorController     = controller.NewAuthDoctorController(authDoctorService, jwtService)
	doctorController         controller.DoctorController         = controller.NewDoctorController(doctorService, jwtService)
	authPharmacistController controller.AuthPharmacistController = controller.NewAuthPharmacistController(authPharmacistService, jwtService)
	pharmacistController     controller.PharmacistController     = controller.NewPharmacistController(pharmacistService, jwtService)
	orderController          controller.OrderController          = controller.NewOrderController(orderService, jwtService)
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Use(CORSMiddleware())
	authRoutes := r.Group("api/admin/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	userRoutes := r.Group("api/admin", middleware.AuthorizeJWT(jwtService))
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
	pharmRoutes := r.Group("api/image")
	{
		pharmRoutes.POST("/", pharmController.Insert)
	}
	doctorAuthRoutes := r.Group("api/doctor/auth")
	{
		doctorAuthRoutes.POST("/login", authDoctorController.LoginDoctor)
		doctorAuthRoutes.POST("/register", authDoctorController.RegisterDoctor)
	}
	doctorRoutes := r.Group("api/doctor")
	{
		doctorRoutes.GET("/profile", doctorController.ProfileDoctor)
		doctorRoutes.PUT("/update", doctorController.UpdateDoctor)
	}
	pharmacistAuthRoutes := r.Group("api/pharmacist/auth")
	{
		pharmacistAuthRoutes.POST("/login", authPharmacistController.LoginPharmacist)
		pharmacistAuthRoutes.POST("/register", authPharmacistController.RegisterPharmacist)
	}
	pharmacistRoutes := r.Group("api/pharmacist")
	{
		pharmacistRoutes.GET("/profile", pharmacistController.ProfilePharmacist)
		pharmacistRoutes.PUT("/update", pharmacistController.UpdatePharmacist)
	}

	orderRoutes := r.Group("api/order", middleware.AuthorizeJWT(jwtService))
	{
		orderRoutes.POST("/", orderController.Insert)
	}
	r.Run()
}
