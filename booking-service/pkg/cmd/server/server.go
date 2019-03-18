package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/hoangbktech/cinema-bookings/booking-service/pkg/kafka"
	"github.com/hoangbktech/cinema-bookings/booking-service/pkg/protocol/rest"
	"log"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/hoangbktech/cinema-bookings/booking-service/pkg/logger"
	"github.com/hoangbktech/cinema-bookings/booking-service/pkg/protocol/grpc"
	"github.com/hoangbktech/cinema-bookings/booking-service/pkg/service/v1"
	cnm "github.com/hoangbktech/cinema-bookings/cinema-service/pkg/api/v1"
	doc "github.com/hoangbktech/cinema-bookings/document-service/pkg/api/v1"
	mov "github.com/hoangbktech/cinema-bookings/movie-service/pkg/api/v1"
	usr "github.com/hoangbktech/cinema-bookings/user-service/pkg/api/v1"
	grpc2 "google.golang.org/grpc"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	UserServerAddress string
	MovieServerAddress string
	CinemaServerAddress string
	DocumentServerAddress string
	NotificationServerAddress string

	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string

	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string

	JaegerAddr string
	ConsulAddr string

	HTTPPort string

	KafkaHost string
	KafkaTopic string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "", "HTTP port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "",
		"Print time format for logger e.g. 2006-01-02T15:04:05Z07:00")

	flag.StringVar(&cfg.UserServerAddress, "user-service-addr", "localhost:9091", "")
	flag.StringVar(&cfg.CinemaServerAddress, "cinema-service-addr", "localhost:9092", "")
	flag.StringVar(&cfg.MovieServerAddress, "movie-service-addr", "localhost:9093", "")
	flag.StringVar(&cfg.DocumentServerAddress, "doc-service-addr", "localhost:9094", "")
	flag.StringVar(&cfg.NotificationServerAddress, "noti-service-addr", "localhost:9095", "")

	flag.StringVar(&cfg.JaegerAddr, "jaeger-service-addr", "localhost:9092", "")
	flag.StringVar(&cfg.ConsulAddr, "consul-service-addr", "localhost:9093", "")

	flag.StringVar(&cfg.KafkaHost, "kafka-host", "localhost:9092", "")
	flag.StringVar(&cfg.KafkaTopic, "kafka-topic", "notification", "")

	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}

	// initialize logger
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	// UserClient
	userConn, err := grpc2.Dial(cfg.UserServerAddress, grpc2.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer userConn.Close()
	userClient := usr.NewUserServiceClient(userConn)

	// MovieClient
	movieConn, err := grpc2.Dial(cfg.MovieServerAddress, grpc2.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer movieConn.Close()
	movieClient := mov.NewMovieServiceClient(movieConn)

	// CinemaClient
	cinemaConn, err := grpc2.Dial(cfg.CinemaServerAddress, grpc2.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer cinemaConn.Close()
	cinemaClient := cnm.NewCinemaServiceClient(cinemaConn)

	// DocumentClient
	documentConn, err := grpc2.Dial(cfg.DocumentServerAddress, grpc2.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer documentConn.Close()
	documentClient := doc.NewDocumentServiceClient(documentConn)


	// Kafka Producer
	kafkaProducer := &kafka.KafkaProducer{KafkaHost: cfg.KafkaHost, Topic: cfg.KafkaTopic}

	v1API := v1.NewBookingServiceServer(db, userClient, cinemaClient, movieClient, documentClient, kafkaProducer)

	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
