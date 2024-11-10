package ip_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/ip/ip_usecase"
)

type UpdateIPRuleHandlerInterface interface {
	Handle(c *gin.Context)
}

type updateIPRuleHandler struct {
	usecase ip_usecase.UpdateIpRulesByIdUseCase
}

func NewUpdateIPRuleHandler(usecase ip_usecase.UpdateIpRulesByIdUseCase) UpdateIPRuleHandlerInterface {
	return &updateIPRuleHandler{
		usecase: usecase,
	}
}

func (c *updateIPRuleHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(400, gin.H{"error": "Id is required"})
		return
	}

	var inputDto ip_usecase.UpdateIpRulesByIdInputDTO

	if err := ctx.BindJSON(&inputDto); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	inputDto.ID = id

	if err := c.usecase.Execute(inputDto); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "IP Rule updated"})
}
