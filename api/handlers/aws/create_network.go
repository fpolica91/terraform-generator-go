package handlers

import (
	"fmt"
	"generatorv/cmds"
	"generatorv/pkgs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCreateNetwork(c *gin.Context) {

	var data struct {
		Provider    string             `json:"provider"`
		CloudType   string             `json:"type"`
		NetworkUnit []pkgs.NetworkUnit `json:"payload"`
	}
	fmt.Println(data, "data")

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resourceString, err := cmds.CreateNetwork(data.NetworkUnit, data.Provider, data.CloudType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"resourcesString": resourceString})

}
