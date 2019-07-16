package routes

import (
	"controllers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func CreateRoutes(){
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.POST("/upload", controllers.Upload)
		api.GET("/user/:key/:id", controllers.ViewAsUser)
		api.GET("/admin/:key/:id", controllers.ViewAsAdmin)
		api.GET("/uploadwithpath/:path", controllers.UploadWithPath)
	}

	router.Run(":3000" )
}

