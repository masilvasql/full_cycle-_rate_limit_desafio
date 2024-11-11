package token_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/token/token_usecase"
)

type DeleteTokenRuleHandler interface {
	Handle(c *gin.Context)
}

type deleteTokenRuleHandler struct {
	usecase token_usecase.DeleteTokenRulesUseCase
}

func NewDeleteTokenRuleHandler(usecase token_usecase.DeleteTokenRulesUseCase) DeleteTokenRuleHandler {
	return &deleteTokenRuleHandler{
		usecase: usecase,
	}
}

func (d *deleteTokenRuleHandler) Handle(c *gin.Context) {
	id := c.Param("id")

	if err := d.usecase.Execute(id); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Token Rule deleted",
	})
}
