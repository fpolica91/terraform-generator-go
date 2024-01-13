// api/main.go
package main

import (
	"generatorv/api/handlers"
	aws_handlers "generatorv/api/handlers/aws"
	_ "generatorv/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-User-Id"}
	router.Use(cors.New(config))

	amazon_handler := router.Group("/api/v1/aws")
	{
		amazon_handler.POST("/compute/function", aws_handlers.HandleCreateCloudFunction)
		amazon_handler.POST("/compute/virtual_machine", aws_handlers.HandleCreateVirtualMachine)
		amazon_handler.POST("/compute/private_cloud", aws_handlers.HandleCreateVirtualPrivateCloud)
		amazon_handler.POST("/provider/create", aws_handlers.HandleCreateProvider)
		amazon_handler.POST("/storage/object_storage", aws_handlers.HandleObjectStorage)
	}

	router.POST("/persist", handlers.HandlePersistState)
	router.Run(":8080") // or another port of your choice
	router.GET("/swagger/*any", gin.WrapH(httpSwagger.WrapHandler))

}
