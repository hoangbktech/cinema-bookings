package v1

import (
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hoangbktech/cinema-bookings/cinema-service/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// cinemaServiceServer is implementation of v1.CinemaServiceServer proto interface
type cinemaServiceServer struct {
	db *sql.DB
}

// NewCinemaServiceServer creates Cinema service
func NewCinemaServiceServer(db *sql.DB) v1.CinemaServiceServer {
	return &cinemaServiceServer{db: db}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *cinemaServiceServer) checkAPI(api string) error {
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
func (s *cinemaServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// Read cinema
func (s *cinemaServiceServer) ReadCinema(ctx context.Context, req *v1.ReadCinemaRequest) (*v1.ReadCinemaResponse, error) {
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

	// query Cinema by ID
	rows, err := c.QueryContext(ctx, "SELECT `id`, `name`, `capacity` FROM cinemas WHERE `id`=?",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Cinema-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Cinema-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Cinema with ID='%d' is not found",
			req.Id))
	}

	// get Cinema data
	var cn v1.Cinema
	if err := rows.Scan(&cn.Id, &cn.Name, &cn.Capacity); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Cinema row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple Cinema rows with ID='%d'",
			req.Id))
	}

	return &v1.ReadCinemaResponse{
		Api:    apiVersion,
		Cinema: &cn,
	}, nil

}

func (s *cinemaServiceServer) ReadShowing(ctx context.Context, req *v1.ReadShowingRequest) (*v1.ReadShowingResponse, error) {

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

	rows, err := c.QueryContext(ctx, "SELECT `id`, `movie_id`, `cinema_id`, `ticket_amounts` FROM showings WHERE `id`=?",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Showing-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Showing-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Showing with ID='%d' is not found",
			req.Id))
	}

	var sw v1.Showing
	var cinemaId int64
	if err := rows.Scan(&sw.Id, &sw.MovieId, &cinemaId, &sw.TicketAmounts); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Cinema row-> "+err.Error())
	}

	cinema, err := s.ReadCinemaById(ctx, cinemaId)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to get cinema-> "+err.Error())
	}
	sw.Cinema = cinema

	return &v1.ReadShowingResponse{
		Api:     apiVersion,
		Showing: &sw,
	}, nil

}

func (s *cinemaServiceServer) ReadCinemaById(ctx context.Context, id int64) (*v1.Cinema, error) {
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	rows, err := c.QueryContext(ctx, "SELECT `id`, `name`, `capacity` FROM cinemas WHERE `id`=?", id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Cinema-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from Cinema-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Cinema with ID='%d' is not found",
			id))
	}

	var cn v1.Cinema
	if err := rows.Scan(&cn.Id, &cn.Name, &cn.Capacity); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from Cinema row-> "+err.Error())
	}
	return &cn, nil

}

func (s *cinemaServiceServer) SetNotify(ctx context.Context, req *v1.SetNotifyRequest) (*v1.SetNotifyResponse, error) {
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	result, err := c.ExecContext(ctx, "UPDATE showings SET `is_notify` = 1 WHERE `id` = ? AND `is_notify` = 0", &req.ShowingId)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to set notify> "+err.Error())
	}
	rowAffected, err := result.RowsAffected()
	return &v1.SetNotifyResponse{
		Api:    apiVersion,
		Result: rowAffected,
	}, nil
}
