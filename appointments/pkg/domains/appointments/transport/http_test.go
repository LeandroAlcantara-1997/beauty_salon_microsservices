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
			name: "success, decodified find app by id",
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
			name: "fail, cannot decodified find app by id",
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

func Test_decodeUpdateApp(t *testing.T) {
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
			name: "success, decodified new update app",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"PUT",
					"/"+"628ed8e442c5ab8d69b6d4fa",
					strings.NewReader(
						`{
							"user_id": 0,
							"salon_id": 1,
							"appointment_date": "2022-06-23T21:12:02.000000001Z"
						}`,
					),
				),
			},
			init: func(r *stdHTTP.Request) *stdHTTP.Request {
				chiCtx := chi.NewRouteContext()
				chiCtx.URLParams.Add("id", "628ed8e442c5ab8d69b6d4fa")
				return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
			},
			want: model.UpsertAppointment{
				ID:              "628ed8e442c5ab8d69b6d4fa",
				UserID:          0,
				SalonID:         1,
				AppointmentDate: time.Date(2022, time.June, 23, 21, 12, 02, 1, time.UTC),
			},
		},
		{
			name: "fail, cannot decodified new update app",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"PUT",
					"/",
					strings.NewReader(
						`{
							"user_id": 0,
							"salon_id": 1,
							"appointment_date": "2022-06-23T21:12:02.000000001Z"
						}`,
					),
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
			got, err := decodeUpdateApp(tt.args.ctx, r)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_decodeAllApp(t *testing.T) {
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
			name: "success, decodified new get all appointments",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"GET",
					"/",
					strings.NewReader(``),
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeAllApp(tt.args.ctx, tt.args.r)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_decodeAppByUser(t *testing.T) {
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
			name: "success, decodified new find app by user",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"GET",
					"/1",
					strings.NewReader(
						`{}`,
					),
				),
			},
			init: func(r *stdHTTP.Request) *stdHTTP.Request {
				chiCtx := chi.NewRouteContext()
				chiCtx.URLParams.Add("id", "1")
				return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
			},
			want: model.FindAppByUser{
				ID: 1,
			},
		},
		{
			name: "fail, cannot read id in the path",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"GET",
					"/1",
					strings.NewReader(
						`{}`,
					),
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
			got, err := decodeAppByUser(tt.args.ctx, r)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_decodeAppBySalon(t *testing.T) {
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
			name: "success, decodified new find app by salon",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"GET",
					"/1",
					strings.NewReader(
						`{}`,
					),
				),
			},
			init: func(r *stdHTTP.Request) *stdHTTP.Request {
				chiCtx := chi.NewRouteContext()
				chiCtx.URLParams.Add("id", "1")
				return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
			},
			want: model.FindAppBySalon{
				ID: 1,
			},
		},
		{
			name: "fail, cannot read id in the path",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"GET",
					"/1",
					strings.NewReader(
						`{}`,
					),
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
			got, err := decodeAppBySalon(tt.args.ctx, r)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_decodeMakeAppointment(t *testing.T) {
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
			name: "success, decodified new make appointment",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"PUT",
					"/1",
					strings.NewReader(`{}`),
				),
			},
			init: func(r *stdHTTP.Request) *stdHTTP.Request {
				chiCtx := chi.NewRouteContext()
				chiCtx.URLParams.Add("id", "1")
				return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
			},
			want: model.MakeAppointment{ID: 1},
		},
		{
			name: "fail, cannot decodified new make appointment",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"PUT",
					"/",
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
			got, err := decodeMakeAppointment(tt.args.ctx, r)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_decodeDeleteApp(t *testing.T) {
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
			name: "success, decodified new delete appointment",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"DELETE",
					"/628ed8e442c5ab8d69b6d4fa",
					strings.NewReader(`{}`),
				),
			},
			init: func(r *stdHTTP.Request) *stdHTTP.Request {
				chiCtx := chi.NewRouteContext()
				chiCtx.URLParams.Add("id", "628ed8e442c5ab8d69b6d4fa")
				return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
			},
			want: model.DeleteAppointment{ID: "628ed8e442c5ab8d69b6d4fa"},
		},
		{
			name: "fail, cannot decodified new delete appointment",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"DELETE",
					"/",
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
			got, err := decodeDeleteApp(tt.args.ctx, r)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_decodeAvailableApp(t *testing.T) {
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
			name: "success, decodified available appointments",
			args: args{
				ctx: context.Background(),
				r: httptest.NewRequest(
					"GET",
					"/available",
					strings.NewReader(``),
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeAvailableApp(tt.args.ctx, tt.args.r)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}
