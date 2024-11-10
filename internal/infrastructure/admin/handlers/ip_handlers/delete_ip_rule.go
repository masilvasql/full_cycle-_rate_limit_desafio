package ip_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/ip/ip_usecase"
)

type DeleteIPRuleHandler interface {
	Handle(c *gin.Context)
}

type deleteIPRuleHandler struct {
	usecase ip_usecase.DeleteIPRulesUseCase
}

func NewDeleteIPRuleHandler(usecase ip_usecase.DeleteIPRulesUseCase) DeleteIPRuleHandler {
	return &deleteIPRuleHandler{
		usecase: usecase,
	}
}

func (d *deleteIPRuleHandler) Handle(c *gin.Context) {
	id := c.Param("id")

	if err := d.usecase.Execute(id); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "IP Rule deleted",
	})
}
