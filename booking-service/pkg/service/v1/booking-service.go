package v1

import (
	"context"
	"database/sql"
	"errors"
	"github.com/hoangbktech/cinema-bookings/booking-service/pkg/api/v1"
	"github.com/hoangbktech/cinema-bookings/booking-service/pkg/kafka"
	"github.com/hoangbktech/cinema-bookings/booking-service/pkg/logger"
	"github.com/hoangbktech/cinema-bookings/booking-service/pkg/model"
	"github.com/hoangbktech/cinema-bookings/booking-service/pkg/util"
	cnm "github.com/hoangbktech/cinema-bookings/cinema-service/pkg/api/v1"
	doc "github.com/hoangbktech/cinema-bookings/document-service/pkg/api/v1"
	mov "github.com/hoangbktech/cinema-bookings/movie-service/pkg/api/v1"
	usr "github.com/hoangbktech/cinema-bookings/user-service/pkg/api/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sync"
)

var (
	mutex = &sync.Mutex{}
)

const (
	// apiVersion is version of API is provided by server
	apiVersion            = "v1"
	notificationThreshold = 0.25
)

type bookingServiceServer struct {
	db             *sql.DB
	userClient     usr.UserServiceClient
	movieClient    mov.MovieServiceClient
	cinemaClient   cnm.CinemaServiceClient
	documentClient doc.DocumentServiceClient
	kafkaProducer  *kafka.KafkaProducer
}

func NewBookingServiceServer(db *sql.DB, uc usr.UserServiceClient, cc cnm.CinemaServiceClient,
	mc mov.MovieServiceClient, dc doc.DocumentServiceClient,
	producer *kafka.KafkaProducer) v1.BookingServiceServer {
	producer.Init()
	return &bookingServiceServer{db: db, userClient: uc, movieClient: mc, cinemaClient: cc, documentClient: dc, kafkaProducer: producer}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *bookingServiceServer) checkAPI(api string) error {
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
func (s *bookingServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

func (s *bookingServiceServer) CreateBooking(ctx context.Context, req *v1.BookingRequest) (*v1.Ticket, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	userRes := &usr.FindUserByPhoneResponse{}
	if res, err := s.findUserByPhone(ctx, req.PhoneNumber); err != nil || res == nil {
		return nil, errors.New("Cannot retrieve userRes by this phone number : " + req.PhoneNumber)
	} else {
		userRes = res
	}

	showingRes := &cnm.ReadShowingResponse{}
	if res, err := s.findShowingById(ctx, req.ShowingId); err != nil || res == nil {
		return nil, errors.New("Cannot retrieve showingRes by this id : " + string(req.ShowingId))
	} else {
		showingRes = res
	}

	movieRes := &mov.ReadMovieResponse{}
	if res, err := s.findMovieById(ctx, showingRes.Showing.MovieId); err != nil || res == nil {
		return nil, errors.New("Cannot retrieve MOVIE by this id : " + string(showingRes.Showing.MovieId))
	} else {
		movieRes = res
	}

	orderId, err := util.GenOrderId()
	if err != nil {
		err := errors.New("there is an error in generating order id")
		return nil, err
	}

	mutex.Lock()
	defer mutex.Unlock()
	bookedAmount, err := s.getBookedAmount(ctx, req.ShowingId)
	if err != nil {
		log.Fatalf("booking error: %v", err)
	}

	if req.TotalAmount > showingRes.Showing.TicketAmounts-bookedAmount {
		err := errors.New("there are not enough tickets available")
		return nil, err
	}

	booking := model.Booking{ShowingId: req.ShowingId, PhoneNumber: userRes.User.PhoneNumber, Amount: req.TotalAmount, OrderId: orderId.String()}
	result, err := s.saveBooking(ctx, &booking)
	if result == 0 || err != nil {
		return nil, err
	}

	// calculate available tickets percentage and send notification to kafka if it reaches over threshold (< 25%)
	availableTicketPct := float64(showingRes.Showing.TicketAmounts-(bookedAmount+req.TotalAmount)) / float64(showingRes.Showing.TicketAmounts)
	s.sendNotification(ctx, showingRes.Showing.Id, &model.Notification{
		Payload: model.Payload{
			Cinema:         showingRes.Showing.Cinema.Name,
			Movie:          movieRes.Movie.Title,
			AvailableSeats: showingRes.Showing.TicketAmounts - bookedAmount,
			TotalSeats:     showingRes.Showing.TicketAmounts,
			BookingUser: model.User{
				Name:        userRes.User.Name,
				LastName:    userRes.User.LastName,
				PhoneNumber: userRes.User.PhoneNumber,
				Email:       userRes.User.Email,
			},
		},
		Method: model.EMAIL,
		Type:   model.ALERT,
	}, availableTicketPct)

	return &v1.Ticket{
		Api:     apiVersion,
		Cinema:  showingRes.Showing.Cinema.Name,
		Movie:   movieRes.Movie.Title,
		User:    &v1.User{Name: userRes.User.Name, LastName: userRes.User.LastName, Email: userRes.User.Email, PhoneNumber: userRes.User.PhoneNumber},
		OrderId: orderId.String(),
	}, nil
}

func (s *bookingServiceServer) getBookedAmount(ctx context.Context, showingId int64) (int64, error) {

	c, err := s.connect(ctx)
	if err != nil {
		return 0, err
	}
	defer c.Close()

	rows, err := c.QueryContext(ctx, "SELECT SUM(amount) FROM bookings WHERE `showing_id`=?", showingId)
	if err != nil {
		return 0, status.Error(codes.Unknown, "failed to select from Booking-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return 0, status.Error(codes.Unknown, "failed to retrieve data from Booking-> "+err.Error())
		}
		return 0, nil
	}

	var bookedAmount int64
	if err := rows.Scan(&bookedAmount); err != nil {
		return 0, nil
	}
	return bookedAmount, nil
}

func (s *bookingServiceServer) saveBooking(ctx context.Context, booking *model.Booking) (int64, error) {

	c, err := s.connect(ctx)
	if err != nil {
		return 0, err
	}
	defer c.Close()

	result, err := c.ExecContext(ctx, "INSERT INTO bookings(`showing_id`, `amount`, `phone_number`, `order_id`) VALUES(?, ?, ?, ?)",
		&booking.ShowingId, &booking.Amount, &booking.PhoneNumber, &booking.OrderId)
	if err != nil {
		return 0, status.Error(codes.Unknown, "failed to save Booking-> "+err.Error())
	}
	return result.LastInsertId()
}

func (s *bookingServiceServer) sendNotification(ctx context.Context, showingId int64, notification *model.Notification, availablePct float64) {
	if availablePct < notificationThreshold {

		// Call cinema service to update notify
		success := s.updateNotify(ctx, showingId)

		if success > 0 {
			// Send an event to kafka
			go s.kafkaProducer.Publish(notification)
		}

	}
}

func (s *bookingServiceServer) updateNotify(ctx context.Context, showingId int64) int64 {
	result, err := s.cinemaClient.SetNotify(ctx, &cnm.SetNotifyRequest{
		Api:       apiVersion,
		Value:     int64(1),
		ShowingId: showingId,
	})
	if err != nil {
		log.Fatalf("set notify error: %v", err)
	}

	logger.Log.Info("update notify successfully ", zap.Int64("success : ", result.Result))

	return result.Result

}

func (s *bookingServiceServer) findUserByPhone(ctx context.Context, phoneNumber string) (*usr.FindUserByPhoneResponse, error) {
	// find user by phone
	userRes, err := s.userClient.FindUserByPhone(ctx, &usr.FindUserByPhoneRequest{
		Api:         apiVersion,
		PhoneNumber: phoneNumber,
	})
	if err != nil {
		log.Fatalf("user error: %v", err)
		return nil, err
	}
	return userRes, nil
}

func (s *bookingServiceServer) findShowingById(ctx context.Context, showingId int64) (*cnm.ReadShowingResponse, error) {
	showingRes, err := s.cinemaClient.ReadShowing(ctx, &cnm.ReadShowingRequest{
		Api: apiVersion,
		Id:  showingId,
	})

	if err != nil {
		log.Fatalf("showing error: %v", err)
		return nil, err
	}
	return showingRes, nil
}

func (s *bookingServiceServer) findMovieById(ctx context.Context, movieId int64) (*mov.ReadMovieResponse, error) {
	movieRes, err := s.movieClient.Read(ctx, &mov.ReadMovieRequest{
		Api: apiVersion,
		Id:  movieId,
	})

	if err != nil {
		log.Fatalf("showing error: %v", err)
		return nil, err
	}
	return movieRes, nil
}
