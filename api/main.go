// api/main.go
package main

import (
	"generatorv/cmds"
	_ "generatorv/docs"
	"generatorv/pkgs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// handleCreateBuckets godoc
// @Summary Create S3 buckets
// @Description Create buckets based on the provided configuration
// @Tags buckets
// @Accept json
// @Produce json
// @Param buckets body []pkgs.Bucket true "Array of Bucket Configurations"
// @Success 200 {string} string "Buckets created successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 405 {string} string "Method Not Allowed"
// @Failure 500 {string} string "Internal Server Error"
// @Router /createbuckets [post]
func handleCreateBuckets(c *gin.Context) {

	var data struct {
		Buckets []pkgs.Bucket `json:"buckets"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the CreateBuckets function from cmds package
	resourcesString, err := cmds.CreateBuckets(data.Buckets)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"resourcesString": resourcesString})

}

func handleCreateProvider(c *gin.Context) {

	var provider pkgs.Prov
	if err := c.BindJSON(&provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	providerString, err := cmds.CreateProvider(provider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	variablesString, err := cmds.CreateVariables(provider.Provider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"providerString": providerString, "variablesString": variablesString})
}

func handleCreateVirtualPrivateCloud(c *gin.Context) {

	var data struct {
		Vpcs []pkgs.VPC `json:"vpcs"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resourceString, err := cmds.CreateVirtualPrivateCloud(data.Vpcs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"resourceString": resourceString})

	// Call the CreateBuckets function from cmds package

}

func handleCreateCloudFunction(c *gin.Context) {
	var data struct {
		Provider string               `json:"provider"`
		Lambdas  []pkgs.CloudFunction `json:"lambdas"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ResourcesString
	resourcesString, err := cmds.CreateCloudFunction(data.Lambdas, data.Provider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"resourcesString": resourcesString})

}

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))
	router.POST("/cloudfunction", handleCreateCloudFunction)
	router.POST("/createvpcs", handleCreateVirtualPrivateCloud)
	router.POST("/createprovider", handleCreateProvider)
	router.POST("/createbuckets", handleCreateBuckets)
	router.Run(":8080") // or another port of your choice

}
