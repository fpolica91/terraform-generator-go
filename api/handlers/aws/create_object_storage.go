package handlers

import (
	"generatorv/cmds"
	"generatorv/pkgs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleObjectStorage(c *gin.Context) {
	var data struct {
		Provider      string               `json:"provider"`
		StorageType   string               `json:"type"`
		ObjectStorage []pkgs.ObjectStorage `json:"payload"`
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
