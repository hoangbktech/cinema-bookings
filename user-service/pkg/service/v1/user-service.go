package v1

import (
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hoangbktech/cinema-bookings/user-service/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// cinemaServiceServer is implementation of v1.CinemaServiceServer proto interface
type userServiceServer struct {
	db *sql.DB
}

// NewUserServiceServer creates User service
func NewUserServiceServer(db *sql.DB) v1.UserServiceServer {
	return &userServiceServer{db: db}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *userServiceServer) checkAPI(api string) error {
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
func (s *userServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}



func (s *userServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
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

	// query User by ID
	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Name`, `LastName`, `Email`, `PhoneNumber` FROM User WHERE `ID`=?",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from User-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from User-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("User with ID='%d' is not found",
			req.Id))
	}

	// get User data
	var usr v1.User
	if err := rows.Scan(&usr.Id, &usr.Name, &usr.LastName, &usr.Email, &usr.PhoneNumber); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from User row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple User rows with ID='%d'",
			req.Id))
	}

	return &v1.ReadResponse{
		Api:  apiVersion,
		User: &usr,
	}, nil

}

func (s *userServiceServer) FindUserByPhone(ctx context.Context, req *v1.FindUserByPhoneRequest) (*v1.FindUserByPhoneResponse, error) {
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

	// query User by Phone
	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Name`, `LastName`, `Email`, `PhoneNumber` FROM User WHERE `PhoneNumber`=?",
		req.PhoneNumber)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from User-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from User-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("User with ID='%d' is not found",
			req.PhoneNumber))
	}

	// get User data
	var usr v1.User
	if err := rows.Scan(&usr.Id, &usr.Name, &usr.LastName, &usr.Email, &usr.PhoneNumber); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from User row-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple User rows with ID='%d'",
			req.PhoneNumber))
	}

	return &v1.FindUserByPhoneResponse{
		Api:  apiVersion,
		User: &usr,
	}, nil
}
