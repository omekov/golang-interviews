package server

import (
	"context"
	"net/http"
	"time"

	"github.com/omekov/golang-interviews/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, port string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
