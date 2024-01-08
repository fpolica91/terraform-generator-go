package handlers

import (
	"encoding/json"
	"fmt"
	"generatorv/api/database"

	"github.com/gin-gonic/gin"
)

func HandlePersistState(c *gin.Context) {
	userId := c.GetHeader("X-User-Id")

	if userId == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var state json.RawMessage

	err := c.BindJSON(&state)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		fmt.Println(err)
	}

	database.Set(userId, state)
}
