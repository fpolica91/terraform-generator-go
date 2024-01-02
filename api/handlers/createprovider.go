package handlers

import (
	"generatorv/cmds"
	"generatorv/pkgs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCreateProvider(c *gin.Context) {

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
