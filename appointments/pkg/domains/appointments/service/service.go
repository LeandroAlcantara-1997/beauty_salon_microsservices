package service

import (
	"context"

	appErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/repository"
	"github.com/facily-tech/go-core/log"
)

type ServiceInterface interface {
	CreateAppointment(context.Context, model.UpsertAppointment) (*model.AppResponse, error)
	UpdateAppointment(context.Context, model.UpsertAppointment) (*model.AppResponse, error)
	FindAllAppointments(context.Context) ([]model.AppResponse, error)
	FindAppointmentsByID(context.Context, model.FindAppointmentsByIDRequest) (model.Appointment, error)
}

type Service struct {
	repository repository.AppointmentRepositoryI
	log        log.Logger
}

func NewService(l log.Logger, repository repository.AppointmentRepositoryI) (*Service, error) {
	if repository == nil {
		return nil, appErr.ErrEmptyRepository
	}
	return &Service{
		log:        l,
		repository: repository,
	}, nil
}

func (s *Service) CreateAppointment(ctx context.Context, app model.UpsertAppointment) (*model.AppResponse, error) {
	return nil, nil
}

func (s *Service) UpdateAppointment(context.Context, model.UpsertAppointment) (*model.AppResponse, error) {
	return nil, nil
}

func (s *Service) FindAllApointments(ctx context.Context) ([]model.AppResponse, error) {
	return nil, nil
}

func (s *Service) FindAppointmentsByID(context.Context, model.FindAppointmentsByIDRequest) (*model.Appointment, error) {
	return nil, nil
}

func (s *Service) FindAppointmentByUserID(ctx context.Context, id int) ([]model.AppResponse, error) {
	return nil, nil
}

func (s *Service) FindAppointmentBySalonID(ctx context.Context, id int) ([]model.AppResponse, error) {
	return nil, nil
}

func (s *Service) MakeAppointment(ctx context.Context) ([]model.AppResponse, error) {
	return nil, nil
}
