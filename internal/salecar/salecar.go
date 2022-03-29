package salecar

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/omekov/golang-interviews/internal/config"
	"github.com/omekov/golang-interviews/pkg/postgresql"
)

var flagConfig = flag.String("config", "local", "path to the config file")
var flagPort = flag.String("port", "80", "specify the port")

func Run() error {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println(*flagPort)
	cfg, err := config.Get(*flagConfig)
	if err != nil {
		return err
	}

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.Postgres.PostgresqlHost,
		cfg.Postgres.PostgresqlPort,
		cfg.Postgres.PostgresqlUser,
		cfg.Postgres.PostgresqlDbname,
		cfg.Postgres.PostgresqlPassword,
	)

	connectionDelay := 1 * time.Minute
	db, err := postgresql.Connection(
		ctx,
		cfg.Postgres.PgDriver,
		dataSourceName,
		connectionDelay,
	)
	if err != nil {
		return err
	}

	fmt.Println(db.Stats())
	// services
	// handlers := delivery.NewHandler(services)
	// srv := server.NewServer(cfg, "8080", nil)

	// go func() {
	// 	if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
	// 		log.Fatalf("error occurred while running http server: %s\n", err.Error())
	// 	}
	// }()

	return nil
}
