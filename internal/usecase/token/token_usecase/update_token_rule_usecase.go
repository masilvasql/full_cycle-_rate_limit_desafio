package token_usecase

import (
	"context"
	"errors"
	"fmt"
	tokenEntityUpdate "github.com/masilvasql/go-rate-limiter/internal/entity"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/token_repository"
	"time"
)

type UpdateTokenRulesByIdInputDTO struct {
	ID         string `json:"id" validate:"required"`
	Token      string `json:"token" validate:"required,token"`
	MaxRequest int    `json:"max_request" validate:"required,numeric"`
	ExpiresIn  string `json:"expires_in" validate:"required"`
}

type UpdateTokenRulesByIdUseCase interface {
	Execute(dto UpdateTokenRulesByIdInputDTO) error
}

type updateTokenRulesByIdUseCase struct {
	tokenRepository token_repository.TokenRepository
}

func NewUpdateTokenRulesByIdUseCase(tokenRepository token_repository.TokenRepository) UpdateTokenRulesByIdUseCase {
	return &updateTokenRulesByIdUseCase{
		tokenRepository: tokenRepository,
	}
}

func (u *updateTokenRulesByIdUseCase) Execute(dto UpdateTokenRulesByIdInputDTO) error {
	if dto.ID == "" {
		return fmt.Errorf("id is required")
	}

	if dto.MaxRequest == 0 {
		return fmt.Errorf("max request is required")
	}

	if dto.ExpiresIn == "" {
		return fmt.Errorf("expires in is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	found, err := u.tokenRepository.GetById(ctx, dto.ID)
	if err != nil && err.Error() == "token not found" {
		return errors.New("token rule not found")
	}

	if found == nil {
		return errors.New("token rule not found")
	}

	entity := tokenEntityUpdate.TokenEntity{
		ID:         dto.ID,
		Token:      found.Token,
		MaxRequest: dto.MaxRequest,
		ExpiresIn:  dto.ExpiresIn,
	}

	if err := u.tokenRepository.Update(ctx, entity); err != nil {
		return err
	}

	return nil
}
