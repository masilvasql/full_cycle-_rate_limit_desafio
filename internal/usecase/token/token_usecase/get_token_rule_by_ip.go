package token_usecase

import (
	"context"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/token_repository"
	"time"
)

type GetTokenRulesByTokenInputDTO struct {
	Token string `json:"token" validate:"required,token"`
}

type GetTokenRulesByTokenOutputDTO struct {
	ID         string `json:"id"`
	Token      string `json:"token"`
	MaxRequest int    `json:"max_request"`
	ExpiresIn  string `json:"expires_in"`
	CreatedAt  string `json:"created_at"`
}

type GetTokenRulesByTokenUseCase interface {
	Execute(dto GetTokenRulesByTokenInputDTO) (*GetTokenRulesByTokenOutputDTO, error)
}

type getTokenRulesByTokenUseCase struct {
	tokenRepository token_repository.TokenRepository
}

func NewGetTokenRulesByTokenUseCase(tokenRepository token_repository.TokenRepository) GetTokenRulesByTokenUseCase {
	return &getTokenRulesByTokenUseCase{
		tokenRepository: tokenRepository,
	}
}

func (g *getTokenRulesByTokenUseCase) Execute(dto GetTokenRulesByTokenInputDTO) (*GetTokenRulesByTokenOutputDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	tokenEntity, err := g.tokenRepository.GetByToken(ctx, dto.Token)
	if err != nil {
		return nil, err
	}

	return &GetTokenRulesByTokenOutputDTO{
		ID:         tokenEntity.ID,
		Token:      tokenEntity.Token,
		MaxRequest: tokenEntity.MaxRequest,
		ExpiresIn:  tokenEntity.ExpiresIn,
		CreatedAt:  tokenEntity.CreatedAt.String(),
	}, nil
}
