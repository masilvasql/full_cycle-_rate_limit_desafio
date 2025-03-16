package middleware

import (
	"context"
	"fmt"
	"github.com/masilvasql/go-rate-limiter/internal/entity"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/interfaces/ip_interfaces"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/interfaces/reate_limiter_interfaces"
	"github.com/masilvasql/go-rate-limiter/internal/infrastructure/database/interfaces/token_interfaces"
	"sync"
	"time"
)

var ErrNotAuthorized = fmt.Errorf("not authorized")
var LimitReached = fmt.Errorf("you have reached the maximum number of requests or actions allowed within a certain time frame")

type RateLimiterInterface interface {
	CheckRateLimit(ip, token string) error
}

type RateLimiter struct {
	LimitedByIP           bool
	LimitedByToken        bool
	TokenRepository       token_interfaces.TokenRepostitoryTokenInterface
	IpRepository          ip_interfaces.IpRepostitoryIPInterface
	RateLimiterRepository reate_limiter_interfaces.RateLimiterRepositoryInterface
	lock                  sync.Mutex
}

func NewRateLimiter(tokenRepository token_interfaces.TokenRepostitoryTokenInterface,
	ipRepository ip_interfaces.IpRepostitoryIPInterface,
	rateLimiterRepository reate_limiter_interfaces.RateLimiterRepositoryInterface,
	limitedByIP bool, limitedByToken bool) *RateLimiter {
	return &RateLimiter{
		LimitedByIP:           limitedByIP,
		LimitedByToken:        limitedByToken,
		TokenRepository:       tokenRepository,
		IpRepository:          ipRepository,
		RateLimiterRepository: rateLimiterRepository,
	}
}

func (r *RateLimiter) CheckRateLimit(ip, token string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if the request is limited by IP or Token
	if token == "" && r.LimitedByIP {
		ipData, err := r.findIP(ctx, ip)
		if err != nil {
			return err
		}
		return r.AllowIp(ctx, *ipData)
	} else if token != "" && r.LimitedByToken { //if token is not empty and the request is limited by token
		tokenData, err := r.findToken(ctx, token)
		if err != nil {
			return err
		}
		return r.AllowToken(ctx, *tokenData)

	} else {
		return ErrNotAuthorized
	}

}

func (r *RateLimiter) findIP(ctx context.Context, ip string) (*entity.IPEntity, error) {
	ipEntity, err := r.IpRepository.GetKey(ctx, ip)
	if err != nil {
		return nil, err
	}
	return ipEntity, nil
}

func (r *RateLimiter) findToken(ctx context.Context, token string) (*entity.TokenEntity, error) {
	ipEntity, err := r.TokenRepository.GetByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return ipEntity, nil
}

func (r *RateLimiter) AllowIp(ctx context.Context, ipEntity entity.IPEntity) error {

	isBanned, err := r.RateLimiterRepository.FindBanKey(ctx, ipEntity.IP)
	if err != nil {
		return err
	}

	if isBanned {
		return LimitReached
	}

	now := time.Now().UnixNano() / int64(time.Millisecond)
	windowStart := now - 1000

	count, err := r.RateLimiterRepository.GetTotRequestInPeriod(ctx, ipEntity.IP, windowStart)
	if err != nil {
		return err
	}

	if count >= int64(ipEntity.MaxRequest) {
		err = r.RateLimiterRepository.AddBanKey(ctx, ipEntity.IP, ipEntity.ExpiresIn)
		if err != nil {
			return err
		}
		return LimitReached
	}

	err = r.RateLimiterRepository.Create(ctx, ipEntity.IP, now)
	if err != nil {
		return err
	}

	return nil

}

func (r *RateLimiter) AllowToken(ctx context.Context, tokenEntity entity.TokenEntity) error {

	isBanned, err := r.RateLimiterRepository.FindBanKey(ctx, tokenEntity.Token)
	if err != nil {
		return err
	}

	if isBanned {
		return LimitReached
	}

	now := time.Now().UnixNano() / int64(time.Millisecond)
	windowStart := now - 1000

	count, err := r.RateLimiterRepository.GetTotRequestInPeriod(ctx, tokenEntity.Token, windowStart)
	if err != nil {
		return err
	}

	if count >= int64(tokenEntity.MaxRequest) {
		err = r.RateLimiterRepository.AddBanKey(ctx, tokenEntity.Token, tokenEntity.ExpiresIn)
		if err != nil {
			return err
		}
		return LimitReached
	}

	err = r.RateLimiterRepository.Create(ctx, tokenEntity.Token, now)
	if err != nil {
		return err
	}

	return nil
}
