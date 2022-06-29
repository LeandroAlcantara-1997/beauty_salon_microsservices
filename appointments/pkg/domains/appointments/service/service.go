package service

import (
	"context"

	appErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/repository"
	"github.com/facily-tech/go-core/log"
)

//go:generate mockgen -destination service_mock.go -package=service -source=service.go
type AppointmentService interface {
	CreateAppointment(context.Context, model.UpsertAppointment) (*model.AppResponse, error)
	UpdateAppointment(context.Context, model.UpsertAppointment) (*model.AppResponse, error)
	MakeAppointment(context.Context, model.MakeAppointment) (*model.AppResponse, error)
	FindAllAppointments(context.Context) ([]model.AppResponse, error)
	FindAvailableAppointments(context.Context) ([]model.AppResponse, error)
	FindAppByID(context.Context, model.FindAppointmentsByIDRequest) (*model.AppResponse, error)
	FindAppByUserID(context.Context, model.FindAppByUser) ([]model.AppResponse, error)
	FindAppBySalonID(context.Context, model.FindAppBySalon) ([]model.AppResponse, error)
	DeleteApp(context.Context, model.DeleteAppointment) error
}

type Service struct {
	repository repository.AppointmentRepositoryI
	memory     repository.AppointmentMemoryI
	log        log.Logger
}

func NewService(l log.Logger, r repository.AppointmentRepositoryI, m repository.AppointmentMemoryI) (*Service, error) {
	if r == nil {
		return nil, appErr.ErrEmptyRepository
	}

	return &Service{
		log:        l,
		repository: r,
		memory:     m,
	}, nil
}

func (s *Service) CreateAppointment(ctx context.Context, app model.UpsertAppointment) (*model.AppResponse, error) {
	var (
		appPersistence *model.Appointment
		err            error
	)
	if appPersistence, err = s.repository.CreateAppointment(ctx, model.NewAppointment(app)); err != nil {
		return nil, err
	}

	appResponse := model.NewAppResponse(*appPersistence)
	return &appResponse, nil
}

func (s *Service) UpdateAppointment(ctx context.Context, app model.UpsertAppointment) (*model.AppResponse, error) {
	var (
		appUpdate *model.Appointment
		err       error
	)
	if appUpdate, err = s.repository.UpdateAppointment(ctx, model.NewAppointment(app)); err != nil {
		return nil, err
	}
	appReponse := model.NewAppResponse(*appUpdate)
	return &appReponse, nil
}

func (s *Service) FindAllAppointments(ctx context.Context) ([]model.AppResponse, error) {
	findAll, err := s.repository.FindAllAppointments(ctx)
	if err != nil {
		return nil, err
	}
	findAllResponse := model.NewAppResponseSlice(findAll)

	return findAllResponse, nil
}

func (s *Service) FindAvailableAppointments(ctx context.Context) ([]model.AppResponse, error) {
	app, err := s.repository.AvaiableAppointment(ctx)
	if err != nil {
		return nil, err
	}

	avaiableResponse := model.NewAppResponseSlice(app)
	return avaiableResponse, nil
}

func (s *Service) FindAppByID(ctx context.Context, app model.FindAppointmentsByIDRequest) (*model.AppResponse, error) {
	findByID, err := s.memory.FindAppByIDMemory(app.ID)
	if err != nil {
		if findByID, err = s.repository.FindAppointmentByID(ctx, app.ID); err != nil {
			return nil, err
		}
		_ = s.memory.CreateAppMemoryByID(*findByID)

		findByIDResp := model.NewAppResponse(*findByID)
		return &findByIDResp, nil
	}

	findByIDResponse := model.NewAppResponse(*findByID)
	return &findByIDResponse, nil
}

func (s *Service) FindAppByUserID(ctx context.Context, id model.FindAppByUser) ([]model.AppResponse, error) {
	app, err := s.memory.FindAppByUserIDMemory(id.ID)
	if err != nil {
		app, err := s.repository.FindAppointmentByUserID(ctx, id.ID)
		if err != nil {
			return nil, err
		}

		_ = s.memory.CreateAppMemoryByUserID(app)
		appResponse := model.NewAppResponseSlice(app)
		return appResponse, nil
	}
	appResponse := model.NewAppResponseSlice(app)
	return appResponse, nil
}

func (s *Service) FindAppBySalonID(ctx context.Context, id model.FindAppBySalon) ([]model.AppResponse, error) {
	app, err := s.memory.FindAppBySalonIDMemory(id.ID)
	if err != nil {
		app, err := s.repository.FindAppointmentBySalonID(ctx, id.ID)
		if err != nil {
			return nil, err
		}

		_ = s.memory.CreateAppMemoryBySalonID(app)
		appResponse := model.NewAppResponseSlice(app)
		return appResponse, nil
	}
	appResponse := model.NewAppResponseSlice(app)
	return appResponse, nil
}

func (s *Service) MakeAppointment(ctx context.Context, make model.MakeAppointment) (*model.AppResponse, error) {
	app, err := s.repository.MakeAppointment(ctx, make.ID, make.UserID)
	if err != nil {
		return nil, err
	}
	appResponse := model.NewAppResponse(*app)
	return &appResponse, nil
}

func (s *Service) DeleteApp(ctx context.Context, app model.DeleteAppointment) error {
	if err := s.repository.DeleteAppointment(ctx, app.ID); err != nil {
		return err
	}

	return nil
}
