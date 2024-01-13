package handlers

import (
	"generatorv/cmds"
	"generatorv/pkgs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleCreateBuckets godoc
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
func HandleObjectStorage(c *gin.Context) {
	var data struct {
		Provider      string               `json:"provider"`
		StorageType   string               `json:"storage_type"`
		ObjectStorage []pkgs.ObjectStorage `json:"object_storage"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the CreateBuckets function from cmds package
	resourcesString, err := cmds.CreateObjectStorage(data.ObjectStorage, data.Provider, data.StorageType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"resourcesString": resourcesString})

}
