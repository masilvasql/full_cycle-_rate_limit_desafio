package ip_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/masilvasql/go-rate-limiter/internal/entity"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type IPRepository struct {
	redisClient redis.Client
}

func NewIPRepository(redisClient redis.Client) *IPRepository {
	return &IPRepository{
		redisClient: redisClient,
	}
}

func (i *IPRepository) Create(ctx context.Context, ipEntity *entity.IPEntity) error {
	err := i.redisClient.HSet(
		ctx,
		ipEntity.ID,
		"Token", ipEntity.IP,
		"MaxRequest", ipEntity.MaxRequest,
		"ExpiresIn", ipEntity.ExpiresIn,
		"CreatedAt", ipEntity.CreatedAt).Err()

	if err != nil {
		return err
	}

	err = i.redisClient.Set(ctx, "ip:"+ipEntity.IP, ipEntity.ID, 0).Err()

	if err != nil {
		return err
	}

	return nil
}

func (i *IPRepository) GetByIP(ctx context.Context, ip string) (*entity.IPEntity, error) {

	id, err := i.redisClient.Get(ctx, "ip:"+ip).Result()

	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("ip not found")
	}

	if err != nil {
		return nil, err
	}

	return i.GetById(ctx, id)
}

func (i *IPRepository) GetById(ctx context.Context, id string) (*entity.IPEntity, error) {
	data, err := i.redisClient.HGetAll(ctx, id).Result()
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("ip not found")
	}

	maxRequestInt, err := strconv.Atoi(data["MaxRequest"])
	if err != nil {
		return nil, err
	}

	timeConvertido, err := time.Parse(time.RFC3339, data["CreatedAt"])
	if err != nil {
		return nil, err
	}

	return &entity.IPEntity{
		ID:         id,
		IP:         data["Token"],
		MaxRequest: maxRequestInt,
		ExpiresIn:  data["ExpiresIn"],
		CreatedAt:  timeConvertido,
	}, nil
}

func (i *IPRepository) Update(ctx context.Context, ipEntity entity.IPEntity) error {
	err := i.redisClient.HSet(
		ctx,
		ipEntity.ID,
		"Token", ipEntity.IP,
		"MaxRequest", ipEntity.MaxRequest,
		"ExpiresIn", ipEntity.ExpiresIn).Err()

	if err != nil {
		return err
	}

	return nil
}

func (i *IPRepository) Delete(ctx context.Context, id string) error {
	err := i.redisClient.Del(ctx, id).Err()
	if err != nil {
		return err
	}

	return nil
}
