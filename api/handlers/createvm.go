package handlers

import (
	"generatorv/cmds"
	"generatorv/pkgs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCreateVirtualMachine(c *gin.Context) {
	var data struct {
		Provider string                `json:"provider"`
		Vms      []pkgs.VirtualMachine `json:"vms"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resourcesString, err := cmds.CreateVirtualMachine(data.Vms, data.Provider)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"resourcesString": resourcesString})

}
