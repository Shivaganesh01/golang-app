package main

import (
	"goapp-gin/controller"
	"goapp-gin/docs"
	"goapp-gin/dto/api"
	"goapp-gin/middleware"
	"goapp-gin/repository"
	"goapp-gin/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// gindump "github.com/tpkeeper/gin-dump"
)

func setupLoggerOutputToFile() {
	f, _ := os.Create("go-gin-app.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Shivaganesh - Video API"
	docs.SwaggerInfo.Description = "Shivaganesh learnt this by Pragmatic Reviews - Youtube Video API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// setupLoggerOutputToFile() - // Call this function to write logs to a file
	// server := gin.Default()  // This is equivalent to gin.New() and server.Use(gin.Recovery(), gin.Logger())
	server := gin.New()
	// server.Static("/css", "./templates/css")
	// server.LoadHTMLGlob("templates/*.html")
	server.Use(gin.Recovery(), gin.Logger())
	// server.Use(middleware.Logger(), middleware.BasicAuth())
	// server.Use(middleware.Logger())
	// server.Use(gindump.Dump())
	defer videoRepository.CloseDB()

	videoAPI := api.NewVideoAPI(loginController, videoController)
	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", videoAPI.Authenticate)
		}

		videos := apiRoutes.Group("/videos", middleware.AuthorizeJWT())
		{
			videos.GET("", videoAPI.GetVideos)
			videos.POST("", videoAPI.CreateVideo)
			videos.PUT(":id", videoAPI.UpdateVideo)
			videos.DELETE(":id", videoAPI.DeleteVideo)
		}
	}

	// server.POST("/login", func(ctx *gin.Context) {
	// 	tokenString := loginController.Login(ctx)
	// 	if tokenString != "" {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"token": tokenString,
	// 		})
	// 	} else {
	// 		ctx.JSON(http.StatusUnauthorized, nil)
	// 	}
	// })

	// apiRoutes := server.Group("/api", middleware.AuthorizeJWT())
	// {
	// 	apiRoutes.GET("/test", func(ctx *gin.Context) {
	// 		ctx.JSON(200, gin.H{
	// 			"message": "Hello...",
	// 		})
	// 	})

	// 	apiRoutes.GET("/videos", func(ctx *gin.Context) {
	// 		ctx.IndentedJSON(http.StatusOK, videoController.FindAll())
	// 	})

	// 	apiRoutes.POST("/video", func(ctx *gin.Context) {
	// 		err := videoController.Save(ctx)
	// 		if err != nil {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		} else {
	// 			ctx.JSON(http.StatusOK, gin.H{"success": "Successfully saved..."})
	// 		}
	// 	})

	// 	apiRoutes.PUT("/video/:id", func(ctx *gin.Context) {
	// 		err := videoController.Update(ctx)
	// 		if err != nil {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		} else {
	// 			ctx.JSON(http.StatusOK, gin.H{"success": "Successfully updated..."})
	// 		}
	// 	})

	// 	apiRoutes.DELETE("/video/:id", func(ctx *gin.Context) {
	// 		err := videoController.Delete(ctx)
	// 		if err != nil {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		} else {
	// 			ctx.JSON(http.StatusOK, gin.H{"success": "Successfully deleted..."})
	// 		}
	// 	})
	// }

	// viewRoutes := server.Group("/view")
	// {
	// 	viewRoutes.GET("/videos", videoController.ShowAll)
	// }

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
