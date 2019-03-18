package main

import (
	"fmt"
	"github.com/hoangbktech/cinema-bookings/movie-service/pkg/cmd/server"
	"os"

)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
