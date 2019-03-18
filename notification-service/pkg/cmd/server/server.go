package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/hoangbktech/cinema-bookings/notification-service/pkg/kafka"
	"github.com/hoangbktech/cinema-bookings/notification-service/pkg/logger"
	"log"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoangbktech/cinema-bookings/notification-service/pkg/protocol/grpc"
	"github.com/hoangbktech/cinema-bookings/notification-service/pkg/service/v1"
	grpc2 "google.golang.org/grpc"
	doc "github.com/hoangbktech/cinema-bookings/document-service/pkg/api/v1"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	DocumentServerAddress string


	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string

	ZookeeperHost string
	Topic string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")

	flag.IntVar(&cfg.LogLevel, "log-level", 0, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "",
		"Print time format for logger e.g. 2006-01-02T15:04:05Z07:00")
	flag.StringVar(&cfg.ZookeeperHost, "zk-host", "", "")
	flag.StringVar(&cfg.Topic, "topic", "", "")

	flag.StringVar(&cfg.DocumentServerAddress, "doc-service-addr", "localhost:9094", "")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	// initialize logger
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	consumer := kafka.KafkaConsumer{ZookeeperHost:cfg.ZookeeperHost, Topic:cfg.Topic}
	consumer.Init()

	// DocumentClient
	documentConn, err := grpc2.Dial(cfg.DocumentServerAddress, grpc2.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer documentConn.Close()
	documentClient := doc.NewDocumentServiceClient(documentConn)

	v1API := v1.NewNotificationServiceServer(ctx, consumer, documentClient)
	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
