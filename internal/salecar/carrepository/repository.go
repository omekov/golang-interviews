package carrepository

import "github.com/jmoiron/sqlx"

type carRepository struct {
	CarTyper  CarTyper
	CarMarker CarMarker
}

func NewCarRepository(db *sqlx.DB) *carRepository {
	return &carRepository{
		CarTyper:  newCarTypeRepository(db),
		CarMarker: newCarMarkRepository(db),
	}
}
