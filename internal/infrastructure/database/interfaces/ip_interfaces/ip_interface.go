package ip_interfaces

import (
	"context"
	"github.com/masilvasql/go-rate-limiter/internal/entity"
)

type IpRepostitoryIPInterface interface {
	Create(ctx context.Context, ipEntity *entity.IPEntity) error
	GetByIP(ctx context.Context, ip string) (*entity.IPEntity, error)
	GetById(ctx context.Context, id string) (*entity.IPEntity, error)
	Update(ctx context.Context, ipEntity entity.IPEntity) error
	Delete(ctx context.Context, id string) error
}