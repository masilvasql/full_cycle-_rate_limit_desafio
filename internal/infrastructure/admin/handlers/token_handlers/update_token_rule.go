package token_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/token/token_usecase"
)

type UpdateTokenRuleHandlerInterface interface {
	Handle(c *gin.Context)
}

type updateTokenRuleHandler struct {
	usecase token_usecase.UpdateTokenRulesByIdUseCase
}

func NewUpdateTokenRuleHandler(usecase token_usecase.UpdateTokenRulesByIdUseCase) UpdateTokenRuleHandlerInterface {
	return &updateTokenRuleHandler{
		usecase: usecase,
	}
}

func (c *updateTokenRuleHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(400, gin.H{"error": "Id is required"})
		return
	}

	var inputDto token_usecase.UpdateTokenRulesByIdInputDTO

	if err := ctx.BindJSON(&inputDto); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	inputDto.ID = id

	if err := c.usecase.Execute(inputDto); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Token Rule updated"})
}
