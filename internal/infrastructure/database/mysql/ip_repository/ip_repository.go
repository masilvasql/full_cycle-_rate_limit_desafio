package ip_repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/masilvasql/go-rate-limiter/internal/entity"
)

type IPMySqlRepository struct {
	db *sql.DB
}

func NewIPMySqlRepository(db *sql.DB) *IPMySqlRepository {
	return &IPMySqlRepository{db: db}
}

func (i *IPMySqlRepository) Create(ctx context.Context, ipEntity *entity.IPEntity) error {

	return errors.New("not implemented")
}

func (i *IPMySqlRepository) GetKey(ctx context.Context, ip string) (*entity.IPEntity, error) {

	return nil, errors.New("not implemented")
}

func (i *IPMySqlRepository) GetById(ctx context.Context, id string) (*entity.IPEntity, error) {
	return nil, errors.New("not implemented")
}

func (i *IPMySqlRepository) GetAll(ctx context.Context) ([]entity.IPEntity, error) {

	return nil, errors.New("not implemented")
}

func (i *IPMySqlRepository) Update(ctx context.Context, ipEntity entity.IPEntity) error {

	return errors.New("not implemented")
}

func (i *IPMySqlRepository) Delete(ctx context.Context, id string) error {

	return errors.New("not implemented")
}
