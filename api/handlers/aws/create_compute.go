package handlers

import (
	"generatorv/cmds"
	"generatorv/pkgs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCreateCompute(c *gin.Context) {
	var data struct {
		Provider    string         `json:"provider"`
		ComputeType string         `json:"type"`
		Compute     []pkgs.Compute `json:"payload"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ResourcesString
	resourcesString, err := cmds.CreateCompute(data.Compute, data.Provider, data.ComputeType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"resourcesString": resourcesString})

}
