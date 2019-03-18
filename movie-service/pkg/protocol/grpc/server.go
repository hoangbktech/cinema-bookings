package grpc

import (
	"context"
	"github.com/hoangbktech/cinema-bookings/movie-service/pkg/logger"
	"github.com/hoangbktech/cinema-bookings/movie-service/pkg/protocol/grpc/middleware"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"github.com/hoangbktech/cinema-bookings/movie-service/pkg/api/v1"
)

// RunServer runs gRPC service to publish Movie service
func RunServer(ctx context.Context, v1API v1.MovieServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	// register service
	server := grpc.NewServer(opts...)
	v1.RegisterMovieServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	logger.Log.Warn("starting gRPC server...")
	return server.Serve(listen)
}