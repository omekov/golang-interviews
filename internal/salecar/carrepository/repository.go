package carrepository

import "github.com/jmoiron/sqlx"

type CarRepository struct {
	CarType CarType
}

func NewCarRepository(db *sqlx.DB) CarRepository {
	return CarRepository{
		CarType: newCarTypeRepository(db),
	}
}
