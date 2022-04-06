package carrepository

import (
	"github.com/jmoiron/sqlx"
	"github.com/omekov/golang-interviews/internal/domain"
)

type CarRepository struct {
	CarTyper   domain.CarTyper
	CarMarker  domain.CarMarker
	CarModeler domain.CarModeler
}

func NewCarRepository(db *sqlx.DB) *CarRepository {
	return &CarRepository{
		CarTyper:   newCarTypeRepository(db),
		CarMarker:  newCarMarkRepository(db),
		CarModeler: newCarModelRepository(db),
	}
}

type store struct {
}
