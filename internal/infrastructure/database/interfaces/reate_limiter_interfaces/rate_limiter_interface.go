package reate_limiter_interfaces

import (
	"context"
)

type RateLimiterRepositoryInterface interface {
	Create(ctx context.Context, requestsKey string, now int64) error
	FindBanKey(ctx context.Context, tokenIP string) (bool, error)
	GetTotRequestInPeriod(ctx context.Context, tokenIP string, windowStart int64) (int64, error)
	AddBanKey(ctx context.Context, tokenIP string, duration string) error
}
