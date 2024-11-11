package ip_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/ip/ip_usecase"
)

type CreateIpRuleHandlerInterface interface {
	Handle(c *gin.Context)
}

type createIpRuleHandler struct {
	usecase ip_usecase.CreateIpRulesUseCase
}

func NewCreateIpRuleHandler(usecase ip_usecase.CreateIpRulesUseCase) CreateIpRuleHandlerInterface {
	return &createIpRuleHandler{
		usecase: usecase,
	}

}

func (c *createIpRuleHandler) Handle(ctx *gin.Context) {
	var inputDto ip_usecase.CreateIpRulesDTO

	if err := ctx.ShouldBindJSON(&inputDto); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.usecase.Execute(inputDto); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Ip Rule created"})

}
