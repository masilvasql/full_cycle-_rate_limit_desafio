package token_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/masilvasql/go-rate-limiter/internal/entity"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type TokenRepository struct {
	redisClient redis.Client
}

func NewTokenRepository(redisClient redis.Client) *TokenRepository {
	return &TokenRepository{
		redisClient: redisClient,
	}
}

func (i *TokenRepository) Create(ctx context.Context, tokenEntity *entity.TokenEntity) (*entity.TokenEntity, error) {
	err := i.redisClient.HSet(
		ctx,
		tokenEntity.ID,
		"Token", tokenEntity.Token,
		"MaxRequest", tokenEntity.MaxRequest,
		"ExpiresIn", tokenEntity.ExpiresIn,
		"CreatedAt", tokenEntity.CreatedAt).Err()

	if err != nil {
		return nil, err
	}

	err = i.redisClient.Set(ctx, "token:"+tokenEntity.Token, tokenEntity.ID, 0).Err()

	if err != nil {
		return nil, err
	}

	err = i.redisClient.SAdd(ctx, "token_keys", tokenEntity.ID).Err()
	if err != nil {
		return nil, err
	}

	return tokenEntity, nil
}

func (i *TokenRepository) GetByToken(ctx context.Context, token string) (*entity.TokenEntity, error) {

	id, err := i.redisClient.Get(ctx, "token:"+token).Result()

	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("token not found")
	}

	if err != nil {
		return nil, err
	}

	return i.GetById(ctx, id)
}

func (i *TokenRepository) GetById(ctx context.Context, id string) (*entity.TokenEntity, error) {
	data, err := i.redisClient.HGetAll(ctx, id).Result()
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("token not found")
	}

	maxRequestInt, err := strconv.Atoi(data["MaxRequest"])
	if err != nil {
		return nil, err
	}

	timeConvertido, err := time.Parse(time.RFC3339, data["CreatedAt"])
	if err != nil {
		return nil, err
	}

	return &entity.TokenEntity{
		ID:         id,
		Token:      data["Token"],
		MaxRequest: maxRequestInt,
		ExpiresIn:  data["ExpiresIn"],
		CreatedAt:  timeConvertido,
	}, nil
}

func (i *TokenRepository) GetAll(ctx context.Context) ([]entity.TokenEntity, error) {
	keys, err := i.redisClient.SMembers(ctx, "token_keys").Result()
	if err != nil {
		return nil, err
	}

	var tokensEntity []entity.TokenEntity

	for _, key := range keys {
		tokenData, err := i.redisClient.HGetAll(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		if len(tokenData) == 0 {
			continue
		}

		maxRequestInt, err := strconv.Atoi(tokenData["MaxRequest"])
		if err != nil {
			return nil, err
		}

		timeConvertido, err := time.Parse(time.RFC3339, tokenData["CreatedAt"])
		if err != nil {
			return nil, err
		}

		tokens := entity.TokenEntity{
			ID:         key,
			Token:      tokenData["Token"],
			MaxRequest: maxRequestInt,
			ExpiresIn:  tokenData["ExpiresIn"],
			CreatedAt:  timeConvertido,
		}

		tokensEntity = append(tokensEntity, tokens)
	}

	return tokensEntity, nil
}

func (i *TokenRepository) Update(ctx context.Context, tokenEntity entity.TokenEntity) error {
	err := i.redisClient.HSet(
		ctx,
		tokenEntity.ID,
		"Token", tokenEntity.Token,
		"MaxRequest", tokenEntity.MaxRequest,
		"ExpiresIn", tokenEntity.ExpiresIn).Err()

	if err != nil {
		return err
	}

	return nil
}

func (i *TokenRepository) Delete(ctx context.Context, id string) error {
	err := i.redisClient.Del(ctx, id).Err()
	if err != nil {
		return err
	}

	return nil
}
