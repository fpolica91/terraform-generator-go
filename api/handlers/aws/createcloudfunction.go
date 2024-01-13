package handlers

import (
	"generatorv/cmds"
	"generatorv/pkgs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCreateCloudFunction(c *gin.Context) {
	var data struct {
		Provider       string               `json:"provider"`
		ComputeType    string               `json:"compute_type"`
		CloudFunctions []pkgs.CloudFunction `json:"cloud_function"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ResourcesString
	resourcesString, err := cmds.CreateCloudFunction(data.CloudFunctions, data.Provider, data.ComputeType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"resourcesString": resourcesString})

}
