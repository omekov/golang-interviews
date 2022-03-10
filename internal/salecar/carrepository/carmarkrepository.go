package carrepository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/omekov/golang-interviews/internal/domain"
)

type carMarkRepository struct {
	db *sqlx.DB
}

func NewCarMarkRepository(db *sqlx.DB) *carMarkRepository {
	return &carMarkRepository{
		db: db,
	}
}

type CarMarkRepository interface {
	Create(ctx context.Context, carMark *domain.CarMark) error
	GetByID(ctx context.Context, ID uint) error
	GetAll(ctx context.Context, carMark *domain.CarMark) error
	Update(ctx context.Context, carMark *domain.CarMark) error
	Delete(ctx context.Context, carMark *domain.CarMark) error
}

func (r *carMarkRepository) Create(ctx context.Context, carMark *domain.CarMark) error {
	return nil
}
