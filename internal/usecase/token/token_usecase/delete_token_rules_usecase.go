package token_usecase

import (
	"context"
	"errors"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/token_repository"
	"time"
)

type DeleteTokenRulesUseCase interface {
	Execute(id string) error
}

type deleteTokenRulesUseCase struct {
	tokenRepository token_repository.TokenRepository
}

func NewDeleteTokenRulesUseCase(tokenRepository token_repository.TokenRepository) DeleteTokenRulesUseCase {
	return &deleteTokenRulesUseCase{
		tokenRepository: tokenRepository,
	}
}

func (d *deleteTokenRulesUseCase) Execute(id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := d.tokenRepository.GetById(ctx, id)
	if err != nil && err.Error() == "token not found" {
		return errors.New("token rule not found")
	}

	if err := d.tokenRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
