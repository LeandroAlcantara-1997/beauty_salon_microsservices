package appointments

import (
	"context"
	"testing"
	"time"

	appErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var fakeAppResponse = model.AppResponse{
	ID:              "629aac9c363519d9a9615369",
	UserID:          1,
	SalonID:         1,
	AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
}

var fakeUpsert = model.UpsertAppointment{
	ID:              "629aac9c363519d9a9615369",
	UserID:          1,
	SalonID:         1,
	AppointmentDate: time.Date(2022, 05, 12, 18, 30, 25, 12, time.Local),
}

func TestCreateAppointment(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		svc     *service.MockAppointmentServiceI
		request interface{}
		ctx     context.Context
	}
	tests := []struct {
		name     string
		args     args
		init     func(s *service.MockAppointmentServiceI, ctx context.Context)
		response interface{}
		err      error
	}{
		{
			name: "success",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: fakeUpsert,
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().CreateAppointment(context.Background(), fakeUpsert).Return(&fakeAppResponse, nil)
			},
			response: &fakeAppResponse,
		},
		{
			name: "fail, returns error",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: fakeUpsert,
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().CreateAppointment(context.Background(), fakeUpsert).Return(nil, appErr.ErrNew)
			},
			response: nil,
			err:      appErr.ErrNew,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init(tt.args.svc, tt.args.ctx)
			response, err := CreateAppointment(tt.args.svc)(tt.args.ctx, tt.args.request)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.response, response)
		})
	}
}

func TestFindAppointmentByID(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		svc     *service.MockAppointmentServiceI
		request interface{}
		ctx     context.Context
	}
	tests := []struct {
		name     string
		args     args
		init     func(s *service.MockAppointmentServiceI, ctx context.Context)
		response interface{}
		err      error
	}{
		{
			name: "success",
			args: args{
				svc: service.NewMockAppointmentServiceI(ctrl),
				request: model.FindAppointmentsByIDRequest{
					ID: fakeAppResponse.ID,
				},
				ctx: context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().FindAppByID(context.Background(), model.FindAppointmentsByIDRequest{
					ID: fakeAppResponse.ID,
				}).Return(&fakeAppResponse, nil)
			},
			response: &fakeAppResponse,
			err:      nil,
		},
		{
			name: "fail, returns error",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: fakeUpsert,
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().FindAppByID(context.Background(), model.FindAppointmentsByIDRequest{
					ID: fakeAppResponse.ID,
				}).Return(nil, appErr.ErrTypeAssertion)
			},
			response: nil,
			err:      appErr.ErrTypeAssertion,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init(tt.args.svc, tt.args.ctx)
			response, err := FindAppointmentByID(tt.args.svc)(tt.args.ctx, tt.args.request)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.response, response)
		})
	}
}

func TestFindAllAppointment(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		svc     *service.MockAppointmentServiceI
		request interface{}
		ctx     context.Context
	}
	tests := []struct {
		name     string
		args     args
		init     func(s *service.MockAppointmentServiceI, ctx context.Context)
		response interface{}
		err      error
	}{
		{
			name: "success",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: nil,
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().FindAllAppointments(context.Background()).Return([]model.AppResponse{fakeAppResponse}, nil)
			},
			response: []model.AppResponse{fakeAppResponse},
			err:      nil,
		},
		{
			name: "fail, return error",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: nil,
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().FindAllAppointments(context.Background()).Return(nil, appErr.ErrDatabase)
			},
			response: nil,
			err:      appErr.ErrDatabase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init(tt.args.svc, tt.args.ctx)
			response, err := FindAllAppointment(tt.args.svc)(tt.args.ctx, tt.args.request)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.response, response)
		})
	}
}

func TestFindAppointmentByUser(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		svc     *service.MockAppointmentServiceI
		request interface{}
		ctx     context.Context
	}
	tests := []struct {
		name     string
		args     args
		init     func(s *service.MockAppointmentServiceI, ctx context.Context)
		response interface{}
		err      error
	}{
		{
			name: "success",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: model.FindAppByUser{ID: 1},
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().FindAppByUserID(ctx, model.FindAppByUser{ID: 1}).Return([]model.AppResponse{fakeAppResponse}, nil)
			},
			response: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "fail, return error",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: model.FindAppByUser{ID: 1},
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().FindAppByUserID(ctx, model.FindAppByUser{ID: 1}).Return(nil, appErr.ErrDatabase)
			},
			err: appErr.ErrDatabase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init(tt.args.svc, tt.args.ctx)
			response, err := FindAppointmentByUser(tt.args.svc)(tt.args.ctx, tt.args.request)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.response, response)
		})
	}
}

func TestFindAppointmentBySalon(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		svc     *service.MockAppointmentServiceI
		request interface{}
		ctx     context.Context
	}
	tests := []struct {
		name     string
		args     args
		init     func(s *service.MockAppointmentServiceI, ctx context.Context)
		response interface{}
		err      error
	}{
		{
			name: "success",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: model.FindAppBySalon{ID: fakeAppResponse.SalonID},
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().FindAppBySalonID(ctx, model.FindAppBySalon{ID: fakeAppResponse.SalonID}).Return([]model.AppResponse{fakeAppResponse}, nil)
			},
			response: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "fail, return error",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: model.FindAppBySalon{ID: fakeAppResponse.SalonID},
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().FindAppBySalonID(ctx, model.FindAppBySalon{ID: fakeAppResponse.SalonID}).Return(nil, appErr.ErrTypeAssertion)
			},
			err: appErr.ErrTypeAssertion,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init(tt.args.svc, tt.args.ctx)
			response, err := FindAppointmentBySalon(tt.args.svc)(tt.args.ctx, tt.args.request)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.response, response)
		})
	}
}

func TestUpdateAppointmentByUser(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		svc     *service.MockAppointmentServiceI
		request interface{}
		ctx     context.Context
	}
	tests := []struct {
		name     string
		args     args
		init     func(s *service.MockAppointmentServiceI, ctx context.Context)
		response interface{}
		err      error
	}{
		{
			name: "success",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: fakeUpsert,
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().UpdateAppointment(ctx, fakeUpsert).Return(&fakeAppResponse, nil)
			},
			response: &fakeAppResponse,
		},
		{
			name: "fail, return error",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: fakeUpsert,
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().UpdateAppointment(ctx, fakeUpsert).Return(nil, appErr.ErrTypeAssertion)
			},
			err: appErr.ErrTypeAssertion,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init(tt.args.svc, tt.args.ctx)
			response, err := UpdateAppointmentByUser(tt.args.svc)(tt.args.ctx, tt.args.request)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.response, response)
		})
	}
}

func TestMakeAppointmentByUser(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		svc     *service.MockAppointmentServiceI
		request interface{}
		ctx     context.Context
	}
	tests := []struct {
		name     string
		args     args
		init     func(s *service.MockAppointmentServiceI, ctx context.Context)
		response interface{}
		err      error
	}{
		{
			name: "success",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: model.MakeAppointment{ID: fakeUpsert.ID, UserID: fakeAppResponse.UserID},
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().MakeAppointment(ctx, model.MakeAppointment{ID: fakeAppResponse.ID, UserID: fakeAppResponse.UserID}).Return(&fakeAppResponse, nil)
			},
			response: &fakeAppResponse,
		},
		{
			name: "fail, return error",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: model.MakeAppointment{ID: fakeUpsert.ID, UserID: fakeUpsert.UserID},
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().MakeAppointment(ctx, model.MakeAppointment{ID: fakeAppResponse.ID, UserID: fakeAppResponse.UserID}).Return(nil, appErr.ErrDatabase)
			},
			err: appErr.ErrDatabase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init(tt.args.svc, tt.args.ctx)
			response, err := MakeAppointmentByUser(tt.args.svc)(tt.args.ctx, tt.args.request)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.response, response)
		})
	}
}

func TestDeleteAppointment(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		svc     *service.MockAppointmentServiceI
		request interface{}
		ctx     context.Context
	}
	tests := []struct {
		name     string
		args     args
		init     func(s *service.MockAppointmentServiceI, ctx context.Context)
		response interface{}
		err      error
	}{
		{
			name: "success",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: model.DeleteAppointment{ID: fakeUpsert.ID},
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().DeleteApp(ctx, model.DeleteAppointment{ID: fakeAppResponse.ID}).Return(nil)
			},
			response: nil,
		},
		{
			name: "fail, return error",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: model.DeleteAppointment{ID: fakeUpsert.ID},
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().DeleteApp(ctx, model.DeleteAppointment{ID: fakeAppResponse.ID}).Return(appErr.ErrDatabase)
			},
			err: appErr.ErrDatabase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init(tt.args.svc, tt.args.ctx)
			response, err := DeleteAppointment(tt.args.svc)(tt.args.ctx, tt.args.request)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.response, response)
		})
	}
}

func TestAvailableAppointment(t *testing.T) {
	var ctrl = gomock.NewController(t)
	ctrl.Finish()
	type args struct {
		svc     *service.MockAppointmentServiceI
		request interface{}
		ctx     context.Context
	}
	tests := []struct {
		name     string
		args     args
		init     func(s *service.MockAppointmentServiceI, ctx context.Context)
		response interface{}
		err      error
	}{
		{
			name: "success",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: nil,
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().FindAvailableAppointments(ctx).Return([]model.AppResponse{fakeAppResponse}, nil)
			},
			response: []model.AppResponse{fakeAppResponse},
		},
		{
			name: "fail, return error",
			args: args{
				svc:     service.NewMockAppointmentServiceI(ctrl),
				request: nil,
				ctx:     context.Background(),
			},
			init: func(s *service.MockAppointmentServiceI, ctx context.Context) {
				s.EXPECT().FindAvailableAppointments(ctx).Return(nil, appErr.ErrDatabase)
			},
			err: appErr.ErrDatabase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init(tt.args.svc, tt.args.ctx)
			response, err := AvailableAppointment(tt.args.svc)(tt.args.ctx, tt.args.request)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.response, response)
		})
	}
}
