package ip_usecase

import (
	"context"
	"github.com/masilvasql/go-rate-limiter/internal/entity"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/interfaces/ip_interfaces"
	"time"
)

type GetAllIpUseCaseInterface interface {
	Execute() ([]entity.IPEntity, error)
}

type GetAllIPUseCase struct {
	ipRepository ip_interfaces.IpRepostitoryIPInterface
}

func NewGetAllIPUseCase(ipRepository ip_interfaces.IpRepostitoryIPInterface) *GetAllIPUseCase {
	return &GetAllIPUseCase{
		ipRepository: ipRepository,
	}
}

func (i *GetAllIPUseCase) Execute() ([]entity.IPEntity, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	output, err := i.ipRepository.GetAll(ctx)
	if err != nil {
		return []entity.IPEntity{}, err
	}
	return output, nil
}
