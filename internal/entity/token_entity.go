package entity

import (
	"github.com/google/uuid"
	"time"
)

type TokenEntity struct {
	ID         string
	Token      string
	MaxRequest int
	ExpiresIn  string
	CreatedAt  time.Time
}

func CreateNewTokenEntity(maxRequest int, expiresIn string) *TokenEntity {
	return &TokenEntity{
		ID:         uuid.New().String(),
		Token:      uuid.New().String(),
		MaxRequest: maxRequest,
		ExpiresIn:  expiresIn,
		CreatedAt:  time.Now(),
	}
}
