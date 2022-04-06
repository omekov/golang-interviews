package service

import "github.com/omekov/golang-interviews/internal/salecar/carrepository"

type Service struct {
	carrepository *carrepository.CarRepository
}

func NewService(carrepository *carrepository.CarRepository) *Service {
	return &Service{
		carrepository: carrepository,
	}
}
