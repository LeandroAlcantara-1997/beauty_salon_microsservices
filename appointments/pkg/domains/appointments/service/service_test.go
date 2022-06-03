package service

import (
	"testing"

	appErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/repository"
	"github.com/facily-tech/go-core/log"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

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
			name: "fail, service cannot be created",
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
