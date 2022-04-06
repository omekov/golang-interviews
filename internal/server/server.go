package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/omekov/golang-interviews/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Server struct {
	httpServer *http.Server
	grpcServer *grpc.Server
	grpcOption []grpc.ServerOption
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
		grpcServer: grpc.NewServer([]grpc.ServerOption{
			grpc.KeepaliveParams(keepalive.ServerParameters{
				MaxConnectionIdle: 5 * time.Minute,
				Timeout:           15 * time.Second,
				MaxConnectionAge:  5 * time.Minute,
				Time:              120 * time.Minute,
			}),
			// grpc.UnaryInterceptor(im.Logger),
			// grpc.ChainUnaryInterceptor(
			// 	grpc_ctxtags.UnaryServerInterceptor(),
			// 	grpc_prometheus.UnaryServerInterceptor,
			// 	grpcrecovery.UnaryServerInterceptor(),
			// ),
		}),
	}
}

func (s *Server) RunGRPC(lis net.Listener) error {
	return s.grpcServer.Serve(lis)
}

func (s *Server) StopGRPC() {
	s.grpcServer.GracefulStop()
	return
}

func (s *Server) RunHTTP() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) StopHTTP(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
