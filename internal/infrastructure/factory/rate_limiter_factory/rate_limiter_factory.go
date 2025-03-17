package rate_limiter_factory

import (
	"database/sql"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/redis/rate_limiter_repository"
	"github.com/redis/go-redis/v9"
)

func NewCreateReateLimiterRepositoryFactory(driver string, redisDatabase *redis.Client, db *sql.DB) *rate_limiter_repository.RateLimiterRepository {
	switch driver {
	case "redis":
		return rate_limiter_repository.NewRateLimiterRepository(*redisDatabase)
	case "mysql":
		panic("mysql is not implemented yet")
	default:
		panic("driver not found")
	}
}
