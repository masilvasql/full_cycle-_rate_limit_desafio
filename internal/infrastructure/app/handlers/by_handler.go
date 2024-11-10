package handlers

import "github.com/gin-gonic/gin"

func ByHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "By",
	})
}
