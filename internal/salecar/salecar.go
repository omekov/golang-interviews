package salecar

import (
	"context"
	"flag"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/omekov/golang-interviews/internal/config"
	"github.com/omekov/golang-interviews/pkg/postgresql"
)

var flagConfig = flag.String("config", "./configs/local.yml", "path to the config file")
var flagPort = flag.String("port", "80", "specify the port")

func Run() error {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println(flagPort)
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

	postgresql.Connection(
		ctx,
		cfg.Postgres.PgDriver,
		dataSourceName,
	)

	// services
	// handlers := delivery.NewHandler(services)
	// HTTP Server
	// srv := server.NewServer(cfg, "8080", handlers.Init(cfg))

	// go func() {
	// 	if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
	// 		logger.Errorf("error occurred while running http server: %s\n", err.Error())
	// 	}
	// }()

	// logger.Info("Server started")

	return nil
}
