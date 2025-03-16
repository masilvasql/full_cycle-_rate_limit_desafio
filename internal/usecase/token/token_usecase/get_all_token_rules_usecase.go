package token_usecase

import (
	"context"
	"github.com/masilvasql/go-rate-limiter/internal/entity"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/token_repository"
	"time"
)

type GetAllTokenRulesUseCaseInterface interface {
	Execute() ([]entity.TokenEntity, error)
}

type GetAllTokenRulesUseCase struct {
	tokenRepository token_repository.TokenRepository
}

func NewGetAllTokenRulesUseCase(tokenRepository token_repository.TokenRepository) *GetAllTokenRulesUseCase {
	return &GetAllTokenRulesUseCase{
		tokenRepository: tokenRepository,
	}
}

func (g *GetAllTokenRulesUseCase) Execute() ([]entity.TokenEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return g.tokenRepository.GetAll(ctx)
}
