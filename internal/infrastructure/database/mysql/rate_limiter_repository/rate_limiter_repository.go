package rate_limiter_repository

import (
	"context"
	"database/sql"
	"errors"
)

type MySqlRateLimiterRepository struct {
	db *sql.DB
}

func NewMySqlRateLimiterRepository(db *sql.DB) *MySqlRateLimiterRepository {
	return &MySqlRateLimiterRepository{db: db}
}

func (r *MySqlRateLimiterRepository) Create(ctx context.Context, requestsKey string, now int64) error {
	return errors.New("not implemented")
}

func (r *MySqlRateLimiterRepository) FindBanKey(ctx context.Context, key string) (bool, error) {
	return false, errors.New("not implemented")
}

func (r *MySqlRateLimiterRepository) AddBanKey(ctx context.Context, banKey string, expiresIn string) error {
	return errors.New("not implemented")
}

func (r *MySqlRateLimiterRepository) GetTotRequestInPeriod(ctx context.Context, tokenIP string, windowStart int64) (int64, error) {
	return 0, errors.New("not implemented")
}
