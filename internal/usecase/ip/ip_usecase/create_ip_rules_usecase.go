package ip_usecase

import (
	"context"
	"errors"
	"github.com/masilvasql/go-rate-limiter/internal/entity"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/ip_repository"
	"time"
)

type CreateIpRulesDTO struct {
	IP         string `json:"ip" validate:"required,ip"`
	MaxRequest int    `json:"max_request" validate:"required,numeric"`
	ExpiresIn  string `json:"expires_in" validate:"required" ` //required type string because it can be "1h", "1m", "1s"
}

type CreateIpRulesUseCase interface {
	Execute(dto CreateIpRulesDTO) error
}

type createIpRulesUseCase struct {
	ipRepository ip_repository.IPRepository
}

func NewCreateIpRulesUseCase(ipRepository ip_repository.IPRepository) CreateIpRulesUseCase {
	return &createIpRulesUseCase{
		ipRepository: ipRepository,
	}
}

func (c *createIpRulesUseCase) Execute(dto CreateIpRulesDTO) error {
	ipEntity := entity.CreateNewIPEntity(dto.IP, dto.MaxRequest, dto.ExpiresIn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	exists, err := c.ipRepository.GetKey(ctx, dto.IP)
	if err != nil && err.Error() == "ip not found" {
		return c.ipRepository.Create(ctx, ipEntity)
	}

	if exists != nil {
		return errors.New("Ip already exists")
	}

	return err

}
