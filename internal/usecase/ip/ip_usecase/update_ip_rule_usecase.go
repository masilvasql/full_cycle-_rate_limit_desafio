package ip_usecase

import (
	"context"
	"errors"
	"fmt"
	ipEntityUpdate "github.com/masilvasql/go-rate-limiter/internal/entity"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/ip_repository"
	"time"
)

type UpdateIpRulesByIdInputDTO struct {
	ID         string `json:"id" validate:"required"`
	IP         string `json:"ip" validate:"required,ip"`
	MaxRequest int    `json:"max_request" validate:"required,numeric"`
	ExpiresIn  string `json:"expires_in" validate:"required"`
}

type UpdateIpRulesByIdUseCase interface {
	Execute(dto UpdateIpRulesByIdInputDTO) error
}

type updateIpRulesByIdUseCase struct {
	ipRepository ip_repository.IPRepository
}

func NewUpdateIpRulesByIdUseCase(ipRepository ip_repository.IPRepository) UpdateIpRulesByIdUseCase {
	return &updateIpRulesByIdUseCase{
		ipRepository: ipRepository,
	}
}

func (u *updateIpRulesByIdUseCase) Execute(dto UpdateIpRulesByIdInputDTO) error {
	if dto.ID == "" {
		return fmt.Errorf("id is required")
	}

	if dto.IP == "" {
		return fmt.Errorf("ip is required")
	}

	if dto.MaxRequest == 0 {
		return fmt.Errorf("max request is required")
	}

	if dto.ExpiresIn == "" {
		return fmt.Errorf("expires in is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := u.ipRepository.GetById(ctx, dto.ID)
	if err != nil && err.Error() == "ip not found" {
		return errors.New("ip rule not found")
	}

	entity := ipEntityUpdate.IPEntity{
		ID:         dto.ID,
		IP:         dto.IP,
		MaxRequest: dto.MaxRequest,
		ExpiresIn:  dto.ExpiresIn,
	}

	if err := u.ipRepository.Update(ctx, entity); err != nil {
		return err
	}

	return nil
}
