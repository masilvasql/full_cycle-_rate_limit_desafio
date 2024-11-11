package token_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/token/token_usecase"
)

type GetTokenRuleByTokenHandlerInterface interface {
	Handle(c *gin.Context)
}

type getTokenRuleByTokenHandler struct {
	usecase token_usecase.GetTokenRulesByTokenUseCase
}

func NewGetTokenRuleByTokenHandler(usecase token_usecase.GetTokenRulesByTokenUseCase) GetTokenRuleByTokenHandlerInterface {
	return &getTokenRuleByTokenHandler{
		usecase: usecase,
	}
}

func (c *getTokenRuleByTokenHandler) Handle(ctx *gin.Context) {
	token := ctx.Param("token")

	if token == "" {
		ctx.JSON(400, gin.H{"error": "Token is required"})
		return
	}

	tokenInputDTO := token_usecase.GetTokenRulesByTokenInputDTO{
		Token: token,
	}

	tokenRule, err := c.usecase.Execute(tokenInputDTO)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, tokenRule)
}
