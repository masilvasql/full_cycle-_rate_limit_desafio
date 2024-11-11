package token_usecase

import (
	"context"
	"github.com/masilvasql/go-rate-limiter/internal/entity"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/token_repository"
	"time"
)

type CreateTokenRulesDTO struct {
	Token      string `json:"token" validate:"required,token"`
	MaxRequest int    `json:"max_request" validate:"required,numeric"`
	ExpiresIn  string `json:"expires_in" validate:"required" `
}

type CreateTokenRulesUseCase interface {
	Execute(dto CreateTokenRulesDTO) (entity.TokenEntity, error)
}

type createTokenRulesUseCase struct {
	tokenRepository token_repository.TokenRepository
}

func NewCreateTokenRulesUseCase(tokenRepository token_repository.TokenRepository) *createTokenRulesUseCase {
	return &createTokenRulesUseCase{
		tokenRepository: tokenRepository,
	}
}

func (c *createTokenRulesUseCase) Execute(dto CreateTokenRulesDTO) (entity.TokenEntity, error) {
	tokenEntity := entity.CreateNewTokenEntity(dto.MaxRequest, dto.ExpiresIn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	output, err := c.tokenRepository.Create(ctx, tokenEntity)
	if err != nil {
		return entity.TokenEntity{}, err
	}

	return *output, nil
}
