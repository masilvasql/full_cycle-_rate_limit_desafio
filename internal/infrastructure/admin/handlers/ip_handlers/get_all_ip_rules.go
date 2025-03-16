package ip_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/ip/ip_usecase"
)

type GetAllIPRulesHandlerInterface interface {
	Handle(c *gin.Context)
}

type GetAllIPRulesHandler struct {
	ipUseCase ip_usecase.GetAllIpUseCaseInterface
}

func NewGetAllIPRulesHandler(ipUseCase ip_usecase.GetAllIpUseCaseInterface) *GetAllIPRulesHandler {
	return &GetAllIPRulesHandler{
		ipUseCase: ipUseCase,
	}
}

func (g *GetAllIPRulesHandler) Handle(c *gin.Context) {
	output, err := g.ipUseCase.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, output)
}
