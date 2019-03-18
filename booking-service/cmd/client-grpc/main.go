package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/hoangbktech/cinema-bookings/booking-service/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewBookingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var id int64 = 2
	var totalAmount int64 = 50
	var phoneNumber = "0934347073"

	// Read
	req2 := v1.BookingRequest{
		Api: apiVersion,
		ShowingId:  id,
		TotalAmount: totalAmount,
		PhoneNumber: phoneNumber,

	}
	res2, err := c.CreateBooking(ctx, &req2)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Printf("Read result: <%+v>\n\n", res2)

}
