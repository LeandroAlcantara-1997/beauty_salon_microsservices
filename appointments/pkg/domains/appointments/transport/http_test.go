package transport

import (
	"context"
	stdHTTP "net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	apErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func Test_decodeFindAppByID(t *testing.T) {
	type args struct {
		ctx context.Context
		r   *stdHTTP.Request
	}
	tests := []struct {
		name string
		args args
		want interface{}
		init func(r *stdHTTP.Request) *stdHTTP.Request
		err  error
	}{
		{
			name: "success, decodified transport",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"GET",
					"/"+"628ed8e442c5ab8d69b6d4fa",
					strings.NewReader(`{}`),
				),
			},
			init: func(r *stdHTTP.Request) *stdHTTP.Request {
				chiCtx := chi.NewRouteContext()
				chiCtx.URLParams.Add("id", "628ed8e442c5ab8d69b6d4fa")
				return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
			},
			want: model.FindAppointmentsByIDRequest{ID: "628ed8e442c5ab8d69b6d4fa"},
		},
		{
			name: "fail, cannot decodified transport",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"GET",
					"/"+"628ed8e442c5ab8d69b6d4fa",
					strings.NewReader(`{}`),
				),
			},
			init: func(r *stdHTTP.Request) *stdHTTP.Request {
				chiCtx := chi.NewRouteContext()
				chiCtx.URLParams.Add("id", "")
				return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
			},
			err: apErr.ErrInvalidPath,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.init(tt.args.r)
			got, err := decodeFindAppByID(tt.args.ctx, r)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_decodeCreateApp(t *testing.T) {
	type args struct {
		ctx context.Context
		r   *stdHTTP.Request
	}
	tests := []struct {
		name string
		args args
		want interface{}
		err  error
	}{
		{
			name: "success, decodified new appointment",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"POST",
					"/",
					strings.NewReader(`{
						"user_id": 0,
						"salon_id": 1,
						"appointment_date": "2022-06-23T21:12:02.000000001Z"
					}`),
				),
			},
			want: model.UpsertAppointment{
				UserID:          0,
				SalonID:         1,
				AppointmentDate: time.Date(2022, time.June, 23, 21, 12, 02, 1, time.UTC),
			},
		},
		{
			name: "fail, cannot decodified new appointment",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"POST",
					"/",
					strings.NewReader(`{
						"user_id": 0,
						"appointment_date": "2022-06-23T21:12:02.000000001Z"
					}`),
				),
			},
			err: apErr.ErrInvalidBody,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeCreateApp(tt.args.ctx, tt.args.r)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}
