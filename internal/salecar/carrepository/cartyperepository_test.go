package carrepository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/omekov/golang-interviews/internal/domain"
	"github.com/omekov/golang-interviews/internal/salecar/carrepository"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestCarTypeRepository(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	carrepository := carrepository.NewCarRepository(db)

	createCarTypeRepository(ctx, t, carrepository)
	getByIDCarTypeRepository(ctx, t, carrepository)
	getAllCarTypeRepository(ctx, t, carrepository)
	updateCarTypeRepository(ctx, t, carrepository)
	deleteCarTypeRepository(ctx, t, carrepository)
}

func createCarTypeRepository(ctx context.Context, t *testing.T, carrepository carrepository.CarRepository) {
	carType := domain.CarType{Name: "test"}
	expandCarType := domain.CarType{ID: 1, Name: "test"}

	t.Run("create", func(t *testing.T) {
		err := carrepository.CarType.Create(ctx, &carType)
		assert.Nil(t, err)
		assert.Equal(t, expandCarType.ID, carType.ID)
		assert.Equal(t, expandCarType.Name, carType.Name)
	})
}

func getByIDCarTypeRepository(ctx context.Context, t *testing.T, carrepository carrepository.CarRepository) {
	var carTypeID uint = 1
	expendCarType := domain.CarType{ID: 1, Name: "test"}
	t.Run("get by id", func(t *testing.T) {
		carType, err := carrepository.CarType.GetByID(ctx, carTypeID)
		if err == sql.ErrNoRows {
			return
		}

		assert.Nil(t, err)
		assert.Equal(t, expendCarType.ID, carType.ID)
		assert.Equal(t, expendCarType.Name, carType.Name)
	})
}
func getAllCarTypeRepository(ctx context.Context, t *testing.T, carrepository carrepository.CarRepository) {
	t.Run("get all", func(t *testing.T) {
		carTypes, err := carrepository.CarType.GetAll(ctx)
		if err == sql.ErrNoRows {
			return
		}

		assert.Nil(t, err)
		assert.Equal(t, 1, len(carTypes))
	})
}

func updateCarTypeRepository(ctx context.Context, t *testing.T, carrepository carrepository.CarRepository) {
	expendCarType := domain.CarType{Name: "test2"}
	carType := domain.CarType{ID: 1, Name: "test2"}
	t.Run("update", func(t *testing.T) {
		err := carrepository.CarType.Update(ctx, &carType)
		if err == sql.ErrNoRows {
			return
		}

		assert.Nil(t, err)
		assert.Equal(t, expendCarType.Name, carType.Name)
	})
}

func deleteCarTypeRepository(ctx context.Context, t *testing.T, carrepository carrepository.CarRepository) {

	carType := domain.CarType{Name: "test"}
	t.Run("delete", func(t *testing.T) {
		err := carrepository.CarType.Create(ctx, &carType)
		assert.Nil(t, err)

		err = carrepository.CarType.Delete(ctx, carType.ID)
		assert.Nil(t, err)

		_, err = carrepository.CarType.GetByID(ctx, carType.ID)
		if err == sql.ErrNoRows {
			return
		}

		assert.Nil(t, err)
	})
}
