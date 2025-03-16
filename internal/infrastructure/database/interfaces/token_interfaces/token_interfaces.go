package token_interfaces

import (
	"context"
	"github.com/masilvasql/go-rate-limiter/internal/entity"
)

type TokenRepostitoryTokenInterface interface {
	Create(ctx context.Context, tokenEntity *entity.TokenEntity) (*entity.TokenEntity, error)
	GetByToken(ctx context.Context, token string) (*entity.TokenEntity, error)
	GetById(ctx context.Context, id string) (*entity.TokenEntity, error)
	GetAll(ctx context.Context) ([]entity.TokenEntity, error)
	Update(ctx context.Context, tokenEntity entity.TokenEntity) error
	Delete(ctx context.Context, id string) error
}
