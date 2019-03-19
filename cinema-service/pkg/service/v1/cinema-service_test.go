package v1

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/hoangbktech/cinema-bookings/cinema-service/pkg/api/v1"
)

func Test_cinemaServiceServer_Read(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewCinemaServiceServer(db)

	type args struct {
		ctx context.Context
		req *v1.ReadCinemaRequest
	}
	tests := []struct {
		name    string
		s       v1.CinemaServiceServer
		args    args
		mock    func()
		want    *v1.ReadCinemaResponse
		wantErr bool
	}{
		{
			name: "OK",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.ReadCinemaRequest{
					Api: "v1",
					Id:  1,
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "capacity"}).
					AddRow(1, "name", 100)
				mock.ExpectQuery("SELECT (.+) FROM cinemas").WithArgs(1).WillReturnRows(rows)
			},
			want: &v1.ReadCinemaResponse{
				Api: "v1",
				Cinema: &v1.Cinema{
					Id:       1,
					Name:     "name",
					Capacity: 100,
				},
			},
		},
		{
			name: "Unsupported API",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.ReadCinemaRequest{
					Api: "v1",
					Id:  1,
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "SELECT failed",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.ReadCinemaRequest{
					Api: "v1",
					Id:  1,
				},
			},
			mock: func() {
				mock.ExpectQuery("SELECT (.+) FROM cinemas").WithArgs(1).
					WillReturnError(errors.New("SELECT failed"))
			},
			wantErr: true,
		},
		{
			name: "Not found",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.ReadCinemaRequest{
					Api: "v1",
					Id:  1,
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "capacity"})
				mock.ExpectQuery("SELECT (.+) FROM Cinema").WithArgs(1).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.ReadCinema(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("cinemaServiceServer.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cinemaServiceServer.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
