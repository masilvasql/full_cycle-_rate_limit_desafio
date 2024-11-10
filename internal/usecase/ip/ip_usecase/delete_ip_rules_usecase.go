package ip_usecase

import (
	"context"
	"errors"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/ip_repository"
	"time"
)

type DeleteIPRulesUseCase interface {
	Execute(id string) error
}

type deleteIPRulesUseCase struct {
	ipRepository ip_repository.IPRepository
}

func NewDeleteIPRulesUseCase(ipRepository ip_repository.IPRepository) DeleteIPRulesUseCase {
	return &deleteIPRulesUseCase{
		ipRepository: ipRepository,
	}
}

func (d *deleteIPRulesUseCase) Execute(id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := d.ipRepository.GetById(ctx, id)
	if err != nil && err.Error() == "ip not found" {
		return errors.New("ip rule not found")
	}

	if err := d.ipRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
