package ip_usecase

import (
	"context"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/ip_repository"
	"time"
)

type GetIpRulesByIpInputDTO struct {
	IP string `json:"ip" validate:"required,ip"`
}

type GetIpRulesByIpOutputDTO struct {
	ID         string `json:"id"`
	IP         string `json:"ip"`
	MaxRequest int    `json:"max_request"`
	ExpiresIn  string `json:"expires_in"`
	CreatedAt  string `json:"created_at"`
}

type GetIpRulesByIpUseCase interface {
	Execute(dto GetIpRulesByIpInputDTO) (*GetIpRulesByIpOutputDTO, error)
}

type getIpRulesByIpUseCase struct {
	ipRepository ip_repository.IPRepository
}

func NewGetIpRulesByIpUseCase(ipRepository ip_repository.IPRepository) GetIpRulesByIpUseCase {
	return &getIpRulesByIpUseCase{
		ipRepository: ipRepository,
	}
}

func (g *getIpRulesByIpUseCase) Execute(dto GetIpRulesByIpInputDTO) (*GetIpRulesByIpOutputDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ipEntity, err := g.ipRepository.GetKey(ctx, dto.IP)
	if err != nil {
		return nil, err
	}

	return &GetIpRulesByIpOutputDTO{
		ID:         ipEntity.ID,
		IP:         ipEntity.IP,
		MaxRequest: ipEntity.MaxRequest,
		ExpiresIn:  ipEntity.ExpiresIn,
		CreatedAt:  ipEntity.CreatedAt.String(),
	}, nil
}
