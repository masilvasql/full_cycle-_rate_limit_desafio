package middleware

import "fmt"

type RateLimiterInterface interface {
	CheckRateLimit(ip, token string)
}

type RateLimiter struct {
	MaximumLimitRequestPerSecond int32
	LimitedByIP                  bool
	LimitedByToken               bool
	ExpiresIn                    string
}

func NewRateLimiter(maximumLimitRequestPerSecond int32, limitedByIP bool, limitedByToken bool, expiresIn string) *RateLimiter {
	return &RateLimiter{
		MaximumLimitRequestPerSecond: maximumLimitRequestPerSecond,
		LimitedByIP:                  limitedByIP,
		LimitedByToken:               limitedByToken,
		ExpiresIn:                    expiresIn,
	}
}

func (r *RateLimiter) CheckRateLimit(ip, token string) {
	fmt.Println("Checking rate limit")
	fmt.Println("Maximum limit request per second: ", r.MaximumLimitRequestPerSecond)
	fmt.Println("Limited by Token: ", r.LimitedByIP)
	fmt.Println("Limited by token: ", r.LimitedByToken)
	fmt.Println("Expires in: ", r.ExpiresIn)

	fmt.Println("Token: ", ip)
	fmt.Println("Token: ", token)
}
