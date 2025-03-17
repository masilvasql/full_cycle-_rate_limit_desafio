package token_repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/masilvasql/go-rate-limiter/internal/entity"
)

type MySqlTokenRepository struct {
	db *sql.DB
}

func NewMySqlTokenRepository(db *sql.DB) *MySqlTokenRepository {
	return &MySqlTokenRepository{db: db}
}

func (i *MySqlTokenRepository) Create(ctx context.Context, tokenEntity *entity.TokenEntity) (*entity.TokenEntity, error) {
	return nil, errors.New("not implemented")
}

func (i *MySqlTokenRepository) GetByToken(ctx context.Context, token string) (*entity.TokenEntity, error) {
	return nil, errors.New("not implemented")
}

func (i *MySqlTokenRepository) GetById(ctx context.Context, id string) (*entity.TokenEntity, error) {
	return nil, errors.New("not implemented")
}

func (i *MySqlTokenRepository) GetAll(ctx context.Context) ([]entity.TokenEntity, error) {
	return nil, errors.New("not implemented")
}

func (i *MySqlTokenRepository) Update(ctx context.Context, tokenEntity entity.TokenEntity) error {
	return errors.New("not implemented")
}

func (i *MySqlTokenRepository) Delete(ctx context.Context, id string) error {
	return errors.New("not implemented")
}
