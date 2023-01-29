package main

import (
	"log"
	"net/http"
	"rest-api/controllers"
	"rest-api/initializers"
	"rest-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
	// AuthController      controllers.AuthController
	// AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	// PostController      controllers.PostController
	// PostRouteController routes.PostRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	// AuthController = controllers.NewAuthController(initializers.DB)
	// AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.UserController{DB: initializers.DB}
	UserRouteController = routes.NewRouteUserController(UserController)

	// PostController = controllers.NewPostController(initializers.DB)
	// PostRouteController = routes.NewRoutePostController(PostController)
	gin.SetMode(gin.DebugMode)
	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/check", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	// AuthRouteController.AuthRoute(router)
	UserRouteController.PostRoute(router)
	// PostRouteController.PostRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
