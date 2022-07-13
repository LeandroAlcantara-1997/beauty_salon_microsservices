package service

import (
	"context"
	"testing"
	"time"

	appErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/log"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/repository"
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
	var l log.AppointmentLogI
	var ctrl *gomock.Controller
	repo := repository.NewMockAppointmentRepositoryI(ctrl)

	srv := Service{repository: repo, log: l}
	type args struct {
		l          log.AppointmentLogI
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
		init func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI)
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
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().CreateAppointment(context.Background(), fakeApp).Return(&fakeApp, nil)
				l := log.NewMockAppointmentLogI(ctrl)
				return repo, l
			},
			want: &fakeAppResponse,
		},
		{
			name: "fail, don't was possible create a new Appointment",
			args: args{
				ctx: context.Background(),
				app: fakeUpsert,
			},
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().CreateAppointment(context.Background(), fakeApp).Return(nil, appErr.ErrDatabase)
				l := log.NewMockAppointmentLogI(ctrl)
				l.EXPECT().LogWithTime(appErr.ErrDatabase).Return(nil)
				return repo, l
			},
			want: nil,
			err:  appErr.ErrDatabase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, l := tt.init()
			s := &Service{
				repository: r,
				memory:     repository.NewMockAppointmentMemoryI(ctrl),
				log:        l,
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
		init func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI)
		want *model.AppResponse
		err  error
	}{
		{
			name: "success, updated Appointment",
			args: args{
				ctx: context.Background(),
				app: fakeUpsert,
			},
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().UpdateAppointment(context.Background(), fakeApp).Return(&fakeApp, nil)
				l := log.NewMockAppointmentLogI(ctrl)
				return repo, l
			},
			want: &fakeAppResponse,
		},
		{
			name: "fail, don't was possible update Appointment",
			args: args{
				ctx: context.Background(),
				app: fakeUpsert,
			},
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().UpdateAppointment(context.Background(), fakeApp).Return(nil, appErr.ErrDatabase)
				l := log.NewMockAppointmentLogI(ctrl)
				l.EXPECT().LogWithTime(appErr.ErrDatabase).Return(nil)
				return repo, l
			},
			err: appErr.ErrDatabase,
		},
		{
			name: "fail, don't found Appointmemnt",
			args: args{
				ctx: context.Background(),
				app: fakeUpsert,
			},
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().UpdateAppointment(context.Background(), fakeApp).Return(nil, appErr.ErrNotFound)
				l := log.NewMockAppointmentLogI(ctrl)
				l.EXPECT().LogWithTime(appErr.ErrNotFound).Return(nil)
				return repo, l
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, l := tt.init()
			s := &Service{
				repository: r,
				memory:     repository.NewMockAppointmentMemoryI(ctrl),
				log:        l,
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
		init func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI)
		args args
		want []model.AppResponse
		err  error
	}{
		{
			name: "success, returned all Appointments",
			args: args{
				ctx: context.Background(),
			},
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAllAppointments(context.Background()).Return([]model.Appointment{
					fakeApp,
				}, nil)
				l := log.NewMockAppointmentLogI(ctrl)
				return repo, l
			},
			want: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "fail, cannot found any Appointments",
			args: args{
				ctx: context.Background(),
			},
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAllAppointments(context.Background()).Return(nil, appErr.ErrNotFound)
				l := log.NewMockAppointmentLogI(ctrl)
				l.EXPECT().LogWithTime(appErr.ErrNotFound).Return(nil)
				return repo, l
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, l := tt.init()
			s := &Service{
				repository: r,
				memory:     repository.NewMockAppointmentMemoryI(ctrl),
				log:        l,
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
		init func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI)
		want []model.AppResponse
		err  error
	}{
		{
			name: "success, found all availables Appointments",
			args: args{
				ctx: context.Background(),
			},
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().AvaiableAppointment(context.Background()).Return([]model.Appointment{
					fakeApp,
				}, nil)
				l := log.NewMockAppointmentLogI(ctrl)
				return repo, l
			},
			want: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "fail, not found any available Appointments",
			args: args{
				ctx: context.Background(),
			},
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().AvaiableAppointment(context.Background()).Return(nil, appErr.ErrNotFound)
				l := log.NewMockAppointmentLogI(ctrl)
				l.EXPECT().LogWithTime(appErr.ErrNotFound).Return(nil)
				return repo, l
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, l := tt.init()
			s := &Service{
				repository: r,
				memory:     repository.NewMockAppointmentMemoryI(ctrl),
				log:        l,
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
		init func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI, *log.MockAppointmentLogI)
		want *model.AppResponse
		err  error
	}{
		{
			name: "success, found Appointment by ID in memory database",
			args: args{
				ctx: context.Background(),
				app: model.FindAppointmentsByIDRequest{ID: fakeApp.ID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI, *log.MockAppointmentLogI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentByID(context.Background(), fakeApp.ID).Return(&fakeApp, nil)
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppByIDMemory(fakeApp.ID).Return(&fakeApp, nil)
				l := log.NewMockAppointmentLogI(ctrl)
				return repo, memory, l
			},
			want: &fakeAppResponse,
		},
		{
			name: "success, found Appointment by ID in database",
			args: args{
				ctx: context.Background(),
				app: model.FindAppointmentsByIDRequest{ID: fakeApp.ID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI, *log.MockAppointmentLogI) {
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppByIDMemory(fakeApp.ID).Return(nil, appErr.ErrMemoryDatabase)
				memory.EXPECT().CreateAppMemoryByID(fakeApp)
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentByID(context.Background(), fakeApp.ID).Return(&fakeApp, nil)
				l := log.NewMockAppointmentLogI(ctrl)
				l.EXPECT().LogWithTime(appErr.ErrMemoryDatabase).Return(nil)
				return repo, memory, l
			},
			want: &fakeAppResponse,
		},
		{
			name: "fail, not found Appointments",
			args: args{
				ctx: context.Background(),
				app: model.FindAppointmentsByIDRequest{ID: fakeApp.ID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI, *log.MockAppointmentLogI) {
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppByIDMemory(fakeApp.ID).Return(nil, appErr.ErrMemoryDatabase)
				memory.EXPECT().CreateAppMemoryByID(fakeApp)
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentByID(context.Background(), fakeApp.ID).Return(nil, appErr.ErrNotFound)
				l := log.NewMockAppointmentLogI(ctrl)
				l.EXPECT().LogWithTime(appErr.ErrNotFound).Return(nil)
				l.EXPECT().LogWithTime(appErr.ErrMemoryDatabase).Return(nil)
				return repo, memory, l
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, m, l := tt.init()
			s := &Service{
				repository: r,
				memory:     m,
				log:        l,
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
		init func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI, *log.MockAppointmentLogI)
		want []model.AppResponse
		err  error
	}{
		{
			name: "success, found Appointment by UserID in memory database",
			args: args{
				ctx: context.Background(),
				id:  model.FindAppByUser{ID: fakeApp.UserID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI, *log.MockAppointmentLogI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentByUserID(context.Background(), fakeApp.UserID).Return([]model.Appointment{fakeApp}, nil)
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppByUserIDMemory(fakeApp.UserID).Return([]model.Appointment{fakeApp}, nil)
				l := log.NewMockAppointmentLogI(ctrl)
				return repo, memory, l
			},
			want: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "success, found Appointment by UserID in database",
			args: args{
				ctx: context.Background(),
				id:  model.FindAppByUser{ID: fakeApp.UserID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI, *log.MockAppointmentLogI) {
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppByUserIDMemory(fakeApp.UserID).Return(nil, appErr.ErrMemoryDatabase)
				memory.EXPECT().CreateAppMemoryByUserID([]model.Appointment{fakeApp}).Return(nil)
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentByUserID(context.Background(), fakeApp.UserID).Return([]model.Appointment{fakeApp}, nil)
				l := log.NewMockAppointmentLogI(ctrl)
				l.EXPECT().LogWithTime(appErr.ErrMemoryDatabase).Return(nil)
				return repo, memory, l
			},
			want: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "fail, don't was possible found Appointment by UserID",
			args: args{
				ctx: context.Background(),
				id:  model.FindAppByUser{ID: fakeApp.UserID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI, *log.MockAppointmentLogI) {
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppByUserIDMemory(fakeApp.UserID).Return(nil, appErr.ErrMemoryDatabase)
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentByUserID(context.Background(), fakeApp.UserID).Return(nil, appErr.ErrNotFound)
				l := log.NewMockAppointmentLogI(ctrl)
				l.EXPECT().LogWithTime(appErr.ErrMemoryDatabase).Return(nil)
				l.EXPECT().LogWithTime(appErr.ErrNotFound).Return(nil)
				return repo, memory, l
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, m, l := tt.init()
			s := &Service{
				repository: r,
				memory:     m,
				log:        l,
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
		init func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI, *log.MockAppointmentLogI)
		want []model.AppResponse
		err  error
	}{
		{
			name: "success, found Appointments by SalonID in database memory",
			args: args{
				ctx: context.Background(),
				id:  model.FindAppBySalon{ID: fakeApp.SalonID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI, *log.MockAppointmentLogI) {
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppBySalonIDMemory(fakeApp.SalonID).Return([]model.Appointment{fakeApp}, nil)
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				l := log.NewMockAppointmentLogI(ctrl)
				return repo, memory, l
			},
			want: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "success, found Appointments by SalonID in database",
			args: args{
				ctx: context.Background(),
				id:  model.FindAppBySalon{ID: fakeApp.SalonID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI, *log.MockAppointmentLogI) {
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppBySalonIDMemory(fakeApp.SalonID).Return(nil, appErr.ErrMemoryDatabase)
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentBySalonID(context.Background(), fakeApp.SalonID).Return([]model.Appointment{fakeApp}, nil)
				memory.EXPECT().CreateAppMemoryBySalonID([]model.Appointment{fakeApp}).Return(nil)
				l := log.NewMockAppointmentLogI(ctrl)
				l.EXPECT().LogWithTime(appErr.ErrMemoryDatabase).Return(nil)
				return repo, memory, l
			},
			want: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "fail, don't was possible found Appointments",
			args: args{
				ctx: context.Background(),
				id:  model.FindAppBySalon{ID: fakeApp.SalonID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *repository.MockAppointmentMemoryI, *log.MockAppointmentLogI) {
				memory := repository.NewMockAppointmentMemoryI(ctrl)
				memory.EXPECT().FindAppBySalonIDMemory(fakeApp.SalonID).Return(nil, appErr.ErrMemoryDatabase)
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().FindAppointmentBySalonID(context.Background(), fakeApp.SalonID).Return(nil, appErr.ErrNotFound)
				l := log.NewMockAppointmentLogI(ctrl)
				l.EXPECT().LogWithTime(appErr.ErrMemoryDatabase).Return(nil)
				l.EXPECT().LogWithTime(appErr.ErrNotFound).Return(nil)
				return repo, memory, l
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, m, l := tt.init()
			s := &Service{
				repository: r,
				memory:     m,
				log:        l,
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
		init func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI)
		want *model.AppResponse
		err  error
	}{
		{
			name: "success, Appointment marked",
			args: args{
				ctx:  context.Background(),
				make: model.MakeAppointment{ID: fakeApp.ID, UserID: fakeApp.UserID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().MakeAppointment(context.Background(), fakeApp.ID, fakeApp.UserID).Return(&fakeApp, nil)
				l := log.NewMockAppointmentLogI(ctrl)
				return repo, l
			},
			want: &fakeAppResponse,
		},
		{
			name: "fail, don't was possible found Appointment",
			args: args{
				ctx:  context.Background(),
				make: model.MakeAppointment{ID: fakeApp.ID, UserID: fakeApp.UserID},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				repo := repository.NewMockAppointmentRepositoryI(ctrl)
				repo.EXPECT().MakeAppointment(context.Background(), fakeApp.ID, fakeApp.UserID).Return(nil, appErr.ErrNotFound)
				l := log.NewMockAppointmentLogI(ctrl)
				l.EXPECT().LogWithTime(appErr.ErrNotFound).Return(nil)
				return repo, l
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, l := tt.init()
			s := &Service{
				repository: r,
				memory:     repository.NewMockAppointmentMemoryI(ctrl),
				log:        l,
			}
			got, err := s.MakeAppointment(tt.args.ctx, tt.args.make)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_DeleteApp(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		ctx context.Context
		app model.DeleteAppointment
	}
	tests := []struct {
		name string
		init func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI)
		args args
		err  error
	}{
		{
			name: "success, deleted appointment",
			args: args{
				ctx: context.Background(),
				app: model.DeleteAppointment{
					ID: fakeApp.ID,
				},
			},
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				r := repository.NewMockAppointmentRepositoryI(ctrl)
				r.EXPECT().DeleteAppointment(context.Background(), fakeApp.ID).Return(nil)
				l := log.NewMockAppointmentLogI(ctrl)
				return r, l
			},
		},
		{
			name: "fail, do not found app for delete",
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				r := repository.NewMockAppointmentRepositoryI(ctrl)
				r.EXPECT().DeleteAppointment(context.Background(), fakeApp.ID).Return(appErr.ErrNotFound)
				l := log.NewMockAppointmentLogI(ctrl)
				l.EXPECT().LogWithTime(appErr.ErrNotFound).Return(nil)
				return r, l
			},
			args: args{
				ctx: context.Background(),
				app: model.DeleteAppointment{
					ID: fakeApp.ID,
				},
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, l := tt.init()
			s := &Service{
				repository: r,
				memory:     repository.NewMockAppointmentMemoryI(ctrl),
				log:        l,
			}
			err := s.DeleteApp(tt.args.ctx, tt.args.app)
			assert.ErrorIs(t, err, tt.err)
		})
	}
}

func TestService_CancelAppointment(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		ctx context.Context
		app model.MakeAppointment
	}
	tests := []struct {
		name string
		init func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI)
		args args
		err  error
	}{
		{
			name: "success, canceled appointment",
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				r := repository.NewMockAppointmentRepositoryI(ctrl)
				r.EXPECT().CancelAppointment(context.Background(), fakeApp.ID, fakeApp.UserID).Return(nil)
				l := log.NewMockAppointmentLogI(ctrl)
				return r, l
			},
			args: args{
				ctx: context.Background(),
				app: model.MakeAppointment{
					ID:     fakeApp.ID,
					UserID: fakeApp.UserID,
				},
			},
		},
		{
			name: "fail, don't possible cancel appointment",
			init: func() (*repository.MockAppointmentRepositoryI, *log.MockAppointmentLogI) {
				r := repository.NewMockAppointmentRepositoryI(ctrl)
				r.EXPECT().CancelAppointment(context.Background(), fakeApp.ID, fakeApp.UserID).Return(appErr.ErrNotFound)
				l := log.NewMockAppointmentLogI(ctrl)
				l.EXPECT().LogWithTime(appErr.ErrNotFound)
				return r, l
			},
			args: args{
				ctx: context.Background(),
				app: model.MakeAppointment{
					ID:     fakeApp.ID,
					UserID: fakeApp.UserID,
				},
			},
			err: appErr.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, l := tt.init()
			s := &Service{
				repository: r,
				memory:     repository.NewMockAppointmentMemoryI(ctrl),
				log:        l,
			}
			err := s.CancelAppointment(tt.args.ctx, tt.args.app)
			assert.ErrorIs(t, err, tt.err)
		})
	}
}
