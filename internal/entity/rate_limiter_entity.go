package entity

import "time"

type RateLimiterEntity struct {
	ID             string
	TokenIP        string
	FirstTime      time.Time
	ExpiredAt      time.Time
	Counter        int32
	IsLimitReached bool
}
