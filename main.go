// api/main.go
package main

import (
	"fmt"
	aws_handlers "generatorv/api/handlers/aws"
	_ "generatorv/docs"
	"os"

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
		amazon_handler.POST("/compute/create", aws_handlers.HandleCreateCompute)
		amazon_handler.POST("/network/create", aws_handlers.HandleCreateNetwork)
		amazon_handler.POST("/provider/create", aws_handlers.HandleCreateProvider)
		amazon_handler.POST("/storage/create", aws_handlers.HandleObjectStorage)
	}
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = "8080" // default port to use if PORT isn't set
	}

	router.Run(":" + port) // or another port of your choice
	router.GET("/swagger/*any", gin.WrapH(httpSwagger.WrapHandler))

}
