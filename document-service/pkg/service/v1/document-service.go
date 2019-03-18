package v1

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hoangbktech/cinema-bookings/document-service/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

type documentServiceServer struct {
	db *sql.DB
}

func NewDocumentServiceServer() v1.DocumentServiceServer {
	return &documentServiceServer{}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *documentServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}


// TO-DO Create HTML Template for Ticket data
func (s *documentServiceServer) CreateTicketEmailTemplate(ctx context.Context, req *v1.Ticket) (*v1.HTMLResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}


	return &v1.HTMLResponse{}, nil
}

// TO-DO Create HTML Template for Alert
func (s *documentServiceServer) CreateAlertEmailTemplate(ctx context.Context, req *v1.AlertTemplateRequest) (*v1.HTMLResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	html := "<p>This is testing template for alert email</p>"

	return &v1.HTMLResponse{Api:apiVersion, Html:html}, nil
}
