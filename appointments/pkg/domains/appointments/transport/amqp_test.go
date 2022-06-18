package transport

import (
	"context"
	"testing"
	"time"

	appErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"github.com/go-kit/kit/transport/amqp"
	delivery "github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func Test_decodeCreateApp(t *testing.T) {
	type args struct {
		ctx context.Context
		r   *delivery.Delivery
	}
	tests := []struct {
		name string
		args args
		want interface{}
		err  error
	}{
		{
			name: "success, decoded new Appoointmet",
			args: args{
				ctx: context.Background(),
				r: &delivery.Delivery{
					Body: []byte(`{
					"user_id": 0,
					"salon_id": 1,
					"appointment_date": "2022-06-23T21:12:02.000000001Z"
				}`),
				},
			},
			want: model.UpsertAppointment{
				UserID:          0,
				SalonID:         1,
				AppointmentDate: time.Date(2022, time.June, 23, 21, 12, 02, 1, time.UTC),
			},
		},
		{
			name: "fail, cannot decoded new Appoointmet",
			args: args{
				ctx: context.Background(),
				r: &delivery.Delivery{
					Body: []byte(`{
					"user_id": 0
					"appointment_date": "2022-06-23T21:12:02.000000001Z"
				}`),
				},
			},
			err: appErr.ErrInvalidBody,
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

func Test_encodeResponseFunc(t *testing.T) {
	type args struct {
		ctx   context.Context
		p     *delivery.Publishing
		input interface{}
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "success, encode response",
			args: args{
				ctx: context.Background(),
				p: &delivery.Publishing{
					Body: []byte(`{
						"user_id": 0,
						"salon_id": 1,
						"appointment_date": "2022-06-23T21:12:02.000000001Z"
					}`),
				},
				input: model.AppResponse{UserID: 0, SalonID: 1, AppointmentDate: time.Date(2022, time.June, 23, 21, 12, 02, 1, time.UTC)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := encodeResponseFunc(tt.args.ctx, tt.args.p, tt.args.input)
			assert.ErrorIs(t, err, tt.err)
		})
	}
}

func Test_errorSubscriber(t *testing.T) {
	type args struct {
		ctx   context.Context
		err   error
		deliv *delivery.Delivery
		ch    amqp.Channel
		p     *delivery.Publishing
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success, error found",
			args: args{
				ctx: context.Background(),
				err: appErr.ErrDatabase,
				deliv: &delivery.Delivery{
					Body: []byte(`{
						"user_id": 0,
						"salon_id": 1,
						"appointment_date": "2022-06-23T21:12:02.000000001Z"
					}`),
				},
				ch: &delivery.Channel{},
				p: &delivery.Publishing{
					Body: []byte(`{
						"user_id": 0,
						"salon_id": 1,
						"appointment_date": "2022-06-23T21:12:02.000000001Z"
					}`),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorSubscriber(tt.args.ctx, tt.args.err, tt.args.deliv, tt.args.ch, tt.args.p)
		})
	}
}
