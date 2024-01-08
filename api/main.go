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
	router.POST("/cloudfunction", handlers.HandleCreateCloudFunction)
	router.POST("/vms", handlers.HandleCreateVirtualMachine)
	router.POST("/createvpcs", handlers.HandleCreateVirtualPrivateCloud)
	router.POST("/createprovider", handlers.HandleCreateProvider)
	router.POST("/createbuckets", handlers.HandleCreateBuckets)
	router.POST("/persist", handlers.HandlePersistState)
	router.Run(":8080") // or another port of your choice
	router.GET("/swagger/*any", gin.WrapH(httpSwagger.WrapHandler))

}
