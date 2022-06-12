package service

import (
	"context"
	"testing"
	"time"

	appErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/repository"
	"github.com/facily-tech/go-core/log"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var fakeUpsert = model.UpsertAppointment{
	ID:              "629aac9c363519d9a9615369",
	UserID:          1,
	SalonID:         1,
	AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
}

var fakeAppResponse = model.AppResponse{
	ID:              "629aac9c363519d9a9615369",
	UserID:          1,
	SalonID:         1,
	AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
}

var fakeApp = model.Appointment{
	ID:              "629aac9c363519d9a9615369",
	UserID:          1,
	SalonID:         1,
	AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
}

func TestNewService(t *testing.T) {
	var l log.Logger
	var ctrl *gomock.Controller
	repo := repository.NewMockAppointmentRepositoryI(ctrl)

	srv := Service{repository: repo, log: l}
	type args struct {
		l          log.Logger
		repository repository.AppointmentRepositoryI
		memory     repository.AppointmentMemoryI
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
			got, err := NewService(tt.args.l, tt.args.repository, tt.args.memory)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_CreateAppointment(t *testing.T) {
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		ctx context.Context
		app model.UpsertAppointment
	}
	tests := []struct {
		name string
		init func() *repository.MockAppointmentRepositoryI
		args args
		want *model.AppResponse
		err  error
	}{
		{
			name: "success, created new Appointment",
			args: args{
				ctx: context.Background(),
				app: fakeUpsert,
			},
			init: func() *repository.MockAppointmentRepositoryI {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().CreateAppointment(context.Background(), fakeApp).Return(&fakeApp, nil)
				return repo
			},
			want: &fakeAppResponse,
		},
		{
			name: "fail, don't was possible create a new Appointment",
			args: args{
				ctx: context.Background(),
				app: fakeUpsert,
			},
			init: func() *repository.MockAppointmentRepositoryI {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().CreateAppointment(context.Background(), fakeApp).Return(nil, appErr.ErrDatabase)
				return repo
			},
			want: nil,
			err:  appErr.ErrDatabase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.init(),
				memory:     repository.NewMockAppointmentMemoryI(ctrl),
				log:        log.NewMockLogger(ctrl),
			}
			got, err := s.CreateAppointment(tt.args.ctx, tt.args.app)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_UpdateAppointment(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		ctx context.Context
		app model.UpsertAppointment
	}
	tests := []struct {
		name string
		args args
		init func() *repository.MockAppointmentRepositoryI
		want *model.AppResponse
		err  error
	}{
		{
			name: "success, updated Appointment",
			args: args{
				ctx: context.Background(),
				app: fakeUpsert,
			},
			init: func() *repository.MockAppointmentRepositoryI {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().UpdateAppointment(context.Background(), fakeApp).Return(&fakeApp, nil)
				return repo
			},
			want: &fakeAppResponse,
		},
		{
			name: "fail, don't was possible update Appointment",
			args: args{
				ctx: context.Background(),
				app: fakeUpsert,
			},
			init: func() *repository.MockAppointmentRepositoryI {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().UpdateAppointment(context.Background(), fakeApp).Return(nil, appErr.ErrDatabase)
				return repo
			},
			err: appErr.ErrDatabase,
		},
		{
			name: "fail, don't found Appointmemnt",
			args: args{
				ctx: context.Background(),
				app: fakeUpsert,
			},
			init: func() *repository.MockAppointmentRepositoryI {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().UpdateAppointment(context.Background(), fakeApp).Return(nil, appErr.ErrNotFound)
				return repo
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.init(),
				memory:     repository.NewMockAppointmentMemoryI(ctrl),
				log:        log.NewMockLogger(ctrl),
			}
			got, err := s.UpdateAppointment(tt.args.ctx, tt.args.app)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_FindAllAppointments(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		init func() *repository.MockAppointmentRepositoryI
		args args
		want []model.AppResponse
		err  error
	}{
		{
			name: "success, returned all Appointments",
			args: args{
				ctx: context.Background(),
			},
			init: func() *repository.MockAppointmentRepositoryI {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAllAppointments(context.Background()).Return([]model.Appointment{
					fakeApp,
				}, nil)
				return repo
			},
			want: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "fail, cannot found any Appointments",
			args: args{
				ctx: context.Background(),
			},
			init: func() *repository.MockAppointmentRepositoryI {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAllAppointments(context.Background()).Return(nil, appErr.ErrNotFound)
				return repo
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.init(),
				memory:     repository.NewMockAppointmentMemoryI(ctrl),
				log:        log.NewMockLogger(ctrl),
			}
			got, err := s.FindAllAppointments(tt.args.ctx)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_FindAvailableAppointments(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		init func() *repository.MockAppointmentRepositoryI
		want []model.AppResponse
		err  error
	}{
		{
			name: "success, found all availables Appointments",
			args: args{
				ctx: context.Background(),
			},
			init: func() *repository.MockAppointmentRepositoryI {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().AvaiableAppointment(context.Background()).Return([]model.Appointment{
					fakeApp,
				}, nil)
				return repo
			},
			want: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "fail, not found any available Appointments",
			args: args{
				ctx: context.Background(),
			},
			init: func() *repository.MockAppointmentRepositoryI {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().AvaiableAppointment(context.Background()).Return(nil, appErr.ErrNotFound)
				return repo
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.init(),
				memory:     repository.NewMockAppointmentMemoryI(ctrl),
				log:        log.NewMockLogger(ctrl),
			}
			got, err := s.FindAvailableAppointments(tt.args.ctx)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_FindAppByID(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		ctx context.Context
		app model.FindAppointmentsByIDRequest
	}
	tests := []struct {
		name string
		args args
		init func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI)
		want *model.AppResponse
		err  error
	}{
		{
			name: "success, found Appointment by ID in memory database",
			args: args{
				ctx: context.Background(),
				app: model.FindAppointmentsByIDRequest{ID: fakeApp.ID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentByID(context.Background(), fakeApp.ID).Return(&fakeApp, nil)
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppByIDMemory(fakeApp.ID).Return(&fakeApp, nil)
				return repo, memory
			},
			want: &fakeAppResponse,
		},
		{
			name: "success, found Appointment by ID in database",
			args: args{
				ctx: context.Background(),
				app: model.FindAppointmentsByIDRequest{ID: fakeApp.ID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI) {
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppByIDMemory(fakeApp.ID).Return(nil, appErr.ErrMemoryDatabase)
				memory.EXPECT().CreateAppMemoryByID(fakeApp)
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentByID(context.Background(), fakeApp.ID).Return(&fakeApp, nil)
				return repo, memory
			},
			want: &fakeAppResponse,
		},
		{
			name: "fail, not found Appointments",
			args: args{
				ctx: context.Background(),
				app: model.FindAppointmentsByIDRequest{ID: fakeApp.ID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI) {
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppByIDMemory(fakeApp.ID).Return(nil, appErr.ErrMemoryDatabase)
				memory.EXPECT().CreateAppMemoryByID(fakeApp)
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentByID(context.Background(), fakeApp.ID).Return(nil, appErr.ErrNotFound)
				return repo, memory
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, m := tt.init()
			s := &Service{
				repository: r,
				memory:     m,
				log:        log.NewMockLogger(ctrl),
			}
			got, err := s.FindAppByID(tt.args.ctx, tt.args.app)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_FindAppByUserID(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		ctx context.Context
		id  model.FindAppByUser
	}
	tests := []struct {
		name string
		args args
		init func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI)
		want []model.AppResponse
		err  error
	}{
		{
			name: "success, found Appointment by UserID in memory database",
			args: args{
				ctx: context.Background(),
				id:  model.FindAppByUser{ID: fakeApp.UserID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentByUserID(context.Background(), fakeApp.UserID).Return([]model.Appointment{fakeApp}, nil)
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppByUserIDMemory(fakeApp.UserID).Return([]model.Appointment{fakeApp}, nil)
				return repo, memory
			},
			want: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "success, found Appointment by UserID in database",
			args: args{
				ctx: context.Background(),
				id:  model.FindAppByUser{ID: fakeApp.UserID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI) {
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppByUserIDMemory(fakeApp.UserID).Return(nil, appErr.ErrMemoryDatabase)
				memory.EXPECT().CreateAppMemoryByUserID([]model.Appointment{fakeApp}).Return(nil)
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentByUserID(context.Background(), fakeApp.UserID).Return([]model.Appointment{fakeApp}, nil)
				return repo, memory
			},
			want: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "fail, don't was possible found Appointment by UserID",
			args: args{
				ctx: context.Background(),
				id:  model.FindAppByUser{ID: fakeApp.UserID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI) {
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppByUserIDMemory(fakeApp.UserID).Return(nil, appErr.ErrMemoryDatabase)
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentByUserID(context.Background(), fakeApp.UserID).Return(nil, appErr.ErrNotFound)
				return repo, memory
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, m := tt.init()
			s := &Service{
				repository: r,
				memory:     m,
				log:        log.NewMockLogger(ctrl),
			}
			got, err := s.FindAppByUserID(tt.args.ctx, tt.args.id)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_FindAppBySalonID(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		ctx context.Context
		id  model.FindAppBySalon
	}
	tests := []struct {
		name string
		args args
		init func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI)
		want []model.AppResponse
		err  error
	}{
		{
			name: "success, found Appointments by SalonID in database memory",
			args: args{
				ctx: context.Background(),
				id:  model.FindAppBySalon{ID: fakeApp.SalonID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI) {
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppBySalonIDMemory(fakeApp.SalonID).Return([]model.Appointment{fakeApp}, nil)
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				return repo, memory
			},
			want: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "success, found Appointments by SalonID in database",
			args: args{
				ctx: context.Background(),
				id:  model.FindAppBySalon{ID: fakeApp.SalonID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI) {
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppBySalonIDMemory(fakeApp.SalonID).Return(nil, appErr.ErrMemoryDatabase)
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentBySalonID(context.Background(), fakeApp.SalonID).Return([]model.Appointment{fakeApp}, nil)
				memory.EXPECT().CreateAppMemoryBySalonID([]model.Appointment{fakeApp}).Return(nil)
				return repo, memory
			},
			want: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "fail, don't was possible found Appointments",
			args: args{
				ctx: context.Background(),
				id:  model.FindAppBySalon{ID: fakeApp.SalonID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI) {
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppBySalonIDMemory(fakeApp.SalonID).Return(nil, appErr.ErrMemoryDatabase)
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentBySalonID(context.Background(), fakeApp.SalonID).Return(nil, appErr.ErrNotFound)
				return repo, memory
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, m := tt.init()
			s := &Service{
				repository: r,
				memory:     m,
				log:        log.NewMockLogger(ctrl),
			}
			got, err := s.FindAppBySalonID(tt.args.ctx, tt.args.id)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_MakeAppointment(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		ctx  context.Context
		make model.MakeAppointment
	}
	tests := []struct {
		name string
		args args
		init func() *repository.MockAppointmentRepositoryI
		want []model.AppResponse
		err  error
	}{
		{
			name: "success, Appointment marked",
			args: args{
				ctx:  context.Background(),
				make: model.MakeAppointment{ID: fakeApp.UserID},
			},
			init: func() *repository.MockAppointmentRepositoryI {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().MakeAppointment(context.Background(), fakeApp.UserID).Return([]model.Appointment{fakeApp}, nil)
				return repo
			},
			want: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "fail, don't was possible found Appointment",
			args: args{
				ctx:  context.Background(),
				make: model.MakeAppointment{ID: fakeApp.UserID},
			},
			init: func() *repository.MockAppointmentRepositoryI {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().MakeAppointment(context.Background(), fakeApp.UserID).Return(nil, appErr.ErrNotFound)
				return repo
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.init()
			s := &Service{
				repository: r,
				memory:     repository.NewMockAppointmentMemoryI(ctrl),
				log:        log.NewMockLogger(ctrl),
			}
			got, err := s.MakeAppointment(tt.args.ctx, tt.args.make)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}
