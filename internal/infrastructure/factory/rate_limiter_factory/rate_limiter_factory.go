package rate_limiter_factory

import (
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/rate_limiter_repository"
	"github.com/redis/go-redis/v9"
)

func NewCreateReateLimiterRepositoryFactory(redisDatabase *redis.Client) *rate_limiter_repository.RateLimiterRepository {
	return rate_limiter_repository.NewRateLimiterRepository(*redisDatabase)
}
