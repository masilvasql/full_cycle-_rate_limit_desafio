package entity

import (
	"github.com/google/uuid"
	"time"
)

type IPEntity struct {
	ID         string
	IP         string
	MaxRequest int
	ExpiresIn  string
	CreatedAt  time.Time
}

func CreateNewIPEntity(ip string, maxRequest int, expiresIn string) *IPEntity {
	return &IPEntity{
		ID:         uuid.New().String(),
		IP:         ip,
		MaxRequest: maxRequest,
		ExpiresIn:  expiresIn,
		CreatedAt:  time.Now(),
	}
}
