package ip_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masilvasql/go-rate-limiter/internal/usecase/ip/ip_usecase"
)

type GetIpRuleByIPHandlerInterface interface {
	Handle(c *gin.Context)
}

type getIpRuleByIPHandler struct {
	usecase ip_usecase.GetIpRulesByIpUseCase
}

func NewGetIpRuleByIPHandler(usecase ip_usecase.GetIpRulesByIpUseCase) GetIpRuleByIPHandlerInterface {
	return &getIpRuleByIPHandler{
		usecase: usecase,
	}
}

func (c *getIpRuleByIPHandler) Handle(ctx *gin.Context) {
	ip := ctx.Param("ip")

	if ip == "" {
		ctx.JSON(400, gin.H{"error": "ip is required"})
		return
	}

	ipInputDTO := ip_usecase.GetIpRulesByIpInputDTO{
		IP: ip,
	}

	ipRule, err := c.usecase.Execute(ipInputDTO)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, ipRule)
}
