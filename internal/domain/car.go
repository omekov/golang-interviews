package domain

import (
	"context"
	"time"
)

type CarType struct {
	ID   uint
	Name string
}

type CarMark struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	CarTypeID uint
	NameRus   string
}

type CarModel struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	CarTypeID uint
	CarMarkID uint
	NameRus   string
}

type CarGeneration struct {
	ID         uint
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CarTypeID  uint
	CarModelID uint
	BeginYear  time.Time
	EndYear    time.Time
}

type CarSerie struct {
	ID              uint
	Name            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CarTypeID       uint
	CarModelID      uint
	CarGenerationID uint
}

type CarModification struct {
	ID         uint
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CarTypeID  uint
	CarSerieID uint
	CarModelID uint
}

type CarOption struct {
	ID          uint
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CarTypeID   uint
	CarOptionID uint
}

type CarOptionValue struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CarTypeID      uint
	IsBase         bool
	CarOptionID    uint
	CarEquipmentID uint
}

type CarEquipment struct {
	ID                uint
	Name              string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CarModificationID uint
	PriceMin          float64
	CarTypeID         uint
	Year              time.Time
}

type CarCharacteristic struct {
	ID                  uint
	Name                string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	CarCharacteristicID uint
	CarTypeID           uint
}

type CarCharacteristicValue struct {
	ID                  uint
	Value               string
	Unit                string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	CarCharacteristicID uint
	CarModificationID   uint
	CarTypeID           uint
}

type CarTyper interface {
	Create(ctx context.Context, carType *CarType) error
	GetByID(ctx context.Context, ID uint) (CarType, error)
	GetAll(ctx context.Context) ([]CarType, error)
	Update(ctx context.Context, carType *CarType) error
	Delete(ctx context.Context, ID uint) error
}

type CarMarker interface {
	Create(ctx context.Context, carMark *CarMark) error
	GetByID(ctx context.Context, ID uint) (CarMark, error)
	GetAll(ctx context.Context) ([]CarMark, error)
	Update(ctx context.Context, carMark *CarMark) error
	Delete(ctx context.Context, ID uint) error
}

type CarModeler interface {
	Create(ctx context.Context, carModel *CarModel) error
	GetByID(ctx context.Context, ID uint) (CarModel, error)
}
