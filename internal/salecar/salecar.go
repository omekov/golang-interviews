package salecar

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	gohttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/omekov/golang-interviews/internal/config"
	"github.com/omekov/golang-interviews/internal/salecar/carrepository"
	"github.com/omekov/golang-interviews/internal/salecar/delivery/http"
	"github.com/omekov/golang-interviews/internal/salecar/service"
	"github.com/omekov/golang-interviews/internal/server"
	"github.com/omekov/golang-interviews/pkg/postgresql"
)

var flagConfig = flag.String("config", "local", "path to the config file")
var flagPort = flag.String("port", "80", "specify the port")

func Run() error {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	// services
	carrepository := carrepository.NewCarRepository(db)
	service := service.NewService(carrepository)
	handlers := http.NewHandler(service)
	srv := server.NewServer(cfg, *flagPort, handlers.Init())

	go func() {
		if err := srv.RunHTTP(); !errors.Is(err, gohttp.ErrServerClosed) {
			log.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	go func() {
		l, err := net.Listen("tcp", "7080")
		if err != nil {

		}
		defer l.Close()

		if err := srv.RunGRPC(l); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Printf("Starting salecar server http://localhost:%s\n", *flagPort)
	fmt.Printf("grpc server http://localhost:%s\n", *flagPort)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	srv.StopHTTP(ctx)
	srv.StopGRPC()
	fmt.Printf("Server Exited Properly\n")

	return nil
}
