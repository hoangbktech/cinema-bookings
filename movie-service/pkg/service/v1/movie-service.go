package v1

import (
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hoangbktech/cinema-bookings/movie-service/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// movieServiceServer is implementation of v1.MovieServiceServer proto interface
type movieServiceServer struct {
	db *sql.DB
}

func NewMovieServiceServer(db *sql.DB) v1.MovieServiceServer {
	return &movieServiceServer{db: db}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *movieServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// connect returns SQL database connection from the pool
func (s *movieServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// Read movie
func (s *movieServiceServer) Read(ctx context.Context, req *v1.ReadMovieRequest) (*v1.ReadMovieResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query Movie by ID
	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description` FROM Movie WHERE `ID`=?",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Movie-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Movie-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Movie with ID='%d' is not found",
			req.Id))
	}

	// get Movie data
	var td v1.Movie
	if err := rows.Scan(&td.Id, &td.Title, &td.Description); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Movie row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ToDo rows with ID='%d'",
			req.Id))
	}

	return &v1.ReadMovieResponse{
		Api:  apiVersion,
		Movie: &td,
	}, nil

}
