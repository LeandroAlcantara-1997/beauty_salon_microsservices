package repository

import (
	"context"

	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
)

//go:generate mockgen -destination repository_mock.go -package=repository -source=repository.go
type AppointmentRepositoryI interface {
	Querier

	Execer
}

type Querier interface {
	FindAllAppointments(context.Context) ([]model.Appointment, error)
	FindAppointmentByID(context.Context, string) (*model.Appointment, error)
	FindAppointmentByUserID(context.Context, int) ([]model.Appointment, error)
	FindAppointmentBySalonID(context.Context, int) ([]model.Appointment, error)
	MakeAppointment(context.Context, int) ([]model.Appointment, error)
	AvaiableAppointment(context.Context) ([]model.Appointment, error)
}

type Execer interface {
	CreateAppointment(context.Context, model.Appointment) (*model.Appointment, error)
	UpdateAppointment(context.Context, model.Appointment) (*model.Appointment, error)
	DeleteAppointment(context.Context, string) error
}

type AppointmentMemoryI interface {
	QuerieMemory

	ExecerMemory
}

type QuerieMemory interface {
	FindAppByIDMemory(string) (*model.Appointment, error)
	FindAppByUserIDMemory(int) ([]model.Appointment, error)
	FindAppBySalonIDMemory(int) ([]model.Appointment, error)
}

type ExecerMemory interface {
	CreateAppMemoryByID(model.Appointment) error
	CreateAppMemoryByUserID([]model.Appointment) error
	CreateAppMemoryBySalonID([]model.Appointment) error
}
