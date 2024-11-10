package token_handlers

import "github.com/gin-gonic/gin"

func CreateTokenRuleHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create IP Rule",
	})
}
