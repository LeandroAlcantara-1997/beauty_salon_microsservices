package service

import (
	"context"
	"testing"
	"time"

	appErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/repository"
	"github.com/facily-tech/go-core/log"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var fakeApp = model.UpsertAppointment{
	ID:              "123",
	UserID:          1,
	SalonID:         1,
	AppointmentDate: time.Date(2022, 12, 24, 22, 30, 20, 10, time.Local),
}

var fakeAppResponse = model.AppResponse{
	ID:              "123",
	UserID:          1,
	SalonID:         1,
	AppointmentDate: time.Date(2022, 12, 24, 22, 30, 20, 10, time.Local),
}

func TestNewService(t *testing.T) {
	var l log.Logger
	var ctrl *gomock.Controller
	repo := repository.NewMockAppointmentRepositoryI(ctrl)

	srv := Service{repository: repo, log: l}
	type args struct {
		l          log.Logger
		repository repository.AppointmentRepositoryI
	}
	tests := []struct {
		name string
		args args
		want *Service
		err  error
	}{
		{
			name: "success, created new service",
			args: args{
				l:          l,
				repository: repo,
			},
			want: &srv,
		},
		{
			name: "fail, service do not created",
			args: args{
				l:          l,
				repository: nil,
			},
			err: appErr.ErrEmptyRepository,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewService(tt.args.l, tt.args.repository)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_CreateAppointment(t *testing.T) {
	var ctrl *gomock.Controller
	repo := repository.NewMockAppointmentRepositoryI(ctrl)
	var l log.Logger
	type fields struct {
		repository repository.AppointmentRepositoryI
		log        log.Logger
	}
	type args struct {
		ctx context.Context
		app model.UpsertAppointment
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.AppResponse
		err    error
	}{
		{
			name: "success, created new appointment",
			fields: fields{
				repository: repo,
				log:        l,
			},
			args: args{
				ctx: context.Background(),
				app: fakeApp,
			},
			want: &fakeAppResponse,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.fields.repository,
				log:        tt.fields.log,
			}
			got, err := s.CreateAppointment(tt.args.ctx, tt.args.app)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}
