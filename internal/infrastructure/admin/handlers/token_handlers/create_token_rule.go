package token_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/token/token_usecase"
)

type CreateTokenRuleHandlerInterface interface {
	Handle(c *gin.Context)
}

type createTokenRuleHandler struct {
	usecase token_usecase.CreateTokenRulesUseCase
}

func NewCreateTokenRuleHandler(usecase token_usecase.CreateTokenRulesUseCase) CreateTokenRuleHandlerInterface {
	return &createTokenRuleHandler{
		usecase: usecase,
	}

}

func (c *createTokenRuleHandler) Handle(ctx *gin.Context) {
	var inputDto token_usecase.CreateTokenRulesDTO

	if err := ctx.ShouldBindJSON(&inputDto); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	output, err := c.usecase.Execute(inputDto)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, output)

}
