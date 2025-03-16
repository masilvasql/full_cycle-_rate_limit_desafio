package rate_limiter_repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RateLimiterRepository struct {
	redisClient redis.Client
}

func NewRateLimiterRepository(redisClient redis.Client) *RateLimiterRepository {
	return &RateLimiterRepository{
		redisClient: redisClient,
	}
}

func (r *RateLimiterRepository) Create(ctx context.Context, requestsKey string, now int64) error {

	keyToCreate := fmt.Sprintf("rate:%s", requestsKey)

	_, err := r.redisClient.ZAdd(ctx, keyToCreate, redis.Z{Score: float64(now), Member: now}).Result()
	if err != nil {
		return err
	}

	err = r.redisClient.Expire(ctx, requestsKey, 10*time.Minute).Err()
	if err != nil {
		log.Println("Erro ao definir TTL:", err)
	}

	return nil
}

func (r *RateLimiterRepository) FindBanKey(ctx context.Context, key string) (bool, error) {
	banKey := fmt.Sprintf("ban:%s", key)
	exists, err := r.redisClient.Exists(ctx, banKey).Result()
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}

func (r *RateLimiterRepository) AddBanKey(ctx context.Context, banKey string, expiresIn string) error {

	duration, err := time.ParseDuration(expiresIn)

	if err != nil {
		return err
	}

	seconds := int32(duration.Seconds())

	banKeyRedis := fmt.Sprintf("ban:%s", banKey)

	err = r.redisClient.SetEx(ctx, banKeyRedis, "1", time.Duration(seconds)*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RateLimiterRepository) GetTotRequestInPeriod(ctx context.Context, tokenIP string, windowStart int64) (int64, error) {
	// Remove todos os elementos cujo score seja menor que windowStart
	requestsKey := fmt.Sprintf("rate:%s", tokenIP)
	_, err := r.redisClient.ZRemRangeByScore(ctx, requestsKey, "0", fmt.Sprintf("%d", windowStart)).Result()
	if err != nil {
		return 0, fmt.Errorf("erro ao remover elementos antigos do ZSET: %w", err)
	}

	// Conta o número de elementos no ZSET (requisições válidas dentro do período)
	count, err := r.redisClient.ZCard(ctx, requestsKey).Result()
	if err != nil {
		return 0, fmt.Errorf("erro ao contar elementos no ZSET: %w", err)
	}

	return count, nil
}
