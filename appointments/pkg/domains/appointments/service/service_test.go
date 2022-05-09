package service

import (
	"testing"

	appErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/repository"
	"github.com/facily-tech/go-core/log"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	l, err := log.NewLoggerZap(log.ZapConfig{})
	assert.NoError(t, err)
	// var mo *mongo.Client
	// repo := repository.NewMongoRepostory(mo, "test", "testCollection")
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
		// {
		// 	name: "success, initialized NewService",
		// 	args: args{
		// 		l:          l,
		// 		repository: repo,
		// 	},
		// 	err: nil,
		// },
		{
			name: "fail, nil repository",
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
