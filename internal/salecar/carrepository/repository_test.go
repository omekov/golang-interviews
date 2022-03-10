package carrepository_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/golang-migrate/migrate/v4"
	dStub "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/omekov/golang-interviews/pkg/postgresql"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	postgresPort := nat.Port("5432/tcp")
	pgContrainer, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image:        "postgres:13-alpine",
				ExposedPorts: []string{postgresPort.Port()},
				Env: map[string]string{
					"POSTGRES_DB":       "postgres",
					"POSTGRES_USER":     "postgres",
					"POSTGRES_PASSWORD": "postgres",
				},
				WaitingFor: wait.ForAll(
					wait.ForLog("database system is ready to accept connections"),
					wait.ForListeningPort(postgresPort),
				),
			},
			Started: true,
		},
	)
	if err != nil {
		log.Fatal(fmt.Errorf("testcontainers.GenericContainer %s", err))
	}

	defer func() {
		if err := pgContrainer.Terminate(ctx); err != nil {
			log.Fatal(fmt.Errorf("pgContrainer.Terminate %s", err))
		}
	}()

	p, err := pgContrainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatal(fmt.Errorf("pgContrainer.MappedPort %s", err))
	}

	log.Println("TestContainer Postgres PORT:", p.Port())
	db, err = postgresql.Connection(ctx, "postgres", fmt.Sprintf("port=%s user=postgres password=postgres dbname=postgres sslmode=disable", p.Port()))
	if err != nil {
		log.Fatal(fmt.Errorf("postgresql.Connection %s", err))
	}

	instance, err := dStub.WithInstance(db.DB, &dStub.Config{})
	if err != nil {
		log.Fatal(fmt.Errorf("dStub.WithInstance %s", err))
	}

	migration, err := migrate.NewWithDatabaseInstance("file://../../../db/migrations/salecar", "postgres", instance)
	if err != nil {
		log.Fatal(fmt.Errorf("migrate.NewWithDatabaseInstance %s", err))
	}

	err = migration.Up()
	if err != nil {
		log.Fatal(fmt.Errorf("migration.Up %s", err))
	}

	os.Exit(m.Run())
}
