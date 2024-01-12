// api/main.go
package main

import (
	"generatorv/api/handlers"
	_ "generatorv/api/handlers"
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
	router.POST("/compute/function", handlers.HandleCreateCloudFunction)
	router.POST("/compute/vm", handlers.HandleCreateVirtualMachine)
	router.POST("/compute/private_cloud", handlers.HandleCreateVirtualPrivateCloud)
	router.POST("/provider/create", handlers.HandleCreateProvider)
	router.POST("/storage/object", handlers.HandleCreateBuckets)
	router.POST("/persist", handlers.HandlePersistState)
	router.Run(":8080") // or another port of your choice
	router.GET("/swagger/*any", gin.WrapH(httpSwagger.WrapHandler))

}
