package v1

import (
"context"
"errors"
"reflect"
"testing"

"gopkg.in/DATA-DOG/go-sqlmock.v1"

"github.com/hoangbktech/cinema-bookings/movie-service/pkg/api/v1"
)

func Test_movieServiceServer_Read(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewMovieServiceServer(db)

	type args struct {
		ctx context.Context
		req *v1.ReadMovieRequest
	}
	tests := []struct {
		name    string
		s       v1.MovieServiceServer
		args    args
		mock    func()
		want    *v1.ReadMovieResponse
		wantErr bool
	}{
		{
			name: "OK",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.ReadMovieRequest{
					Api: "v1",
					Id:  1,
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"ID", "Title", "Description"}).
					AddRow(1, "title", "description")
				mock.ExpectQuery("SELECT (.+) FROM Movie").WithArgs(1).WillReturnRows(rows)
			},
			want: &v1.ReadMovieResponse{
				Api: "v1",
				Movie: &v1.Movie{
					Id:          1,
					Title:       "title",
					Description: "description",
				},
			},
		},
		{
			name: "Unsupported API",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.ReadMovieRequest{
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
				req: &v1.ReadMovieRequest{
					Api: "v1",
					Id:  1,
				},
			},
			mock: func() {
				mock.ExpectQuery("SELECT (.+) FROM Movie").WithArgs(1).
					WillReturnError(errors.New("SELECT failed"))
			},
			wantErr: true,
		},
		{
			name: "Not found",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.ReadMovieRequest{
					Api: "v1",
					Id:  1,
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"ID", "Title", "Description"})
				mock.ExpectQuery("SELECT (.+) FROM Movie").WithArgs(1).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Read(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("movieServiceServer.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("movieServiceServer.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
