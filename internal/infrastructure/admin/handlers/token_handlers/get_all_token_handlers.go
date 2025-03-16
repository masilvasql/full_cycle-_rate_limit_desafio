package token_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/token/token_usecase"
)

type GetAllTokenHandlersInterface interface {
	Handle(g *gin.Context)
}

type GetAllTokenHandlers struct {
	usecase token_usecase.GetAllTokenRulesUseCaseInterface
}

func NewGetAllTokenHandlers(usecase token_usecase.GetAllTokenRulesUseCaseInterface) *GetAllTokenHandlers {
	return &GetAllTokenHandlers{
		usecase: usecase,
	}
}

func (ctx *GetAllTokenHandlers) Handle(g *gin.Context) {
	output, err := ctx.usecase.Execute()
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
	}

	g.JSON(200, output)
}
