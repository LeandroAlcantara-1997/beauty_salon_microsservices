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
	// MakeAppointment(context.Context) ([]model.Appointment, error)
}

type Execer interface {
	CreateAppointment(context.Context, model.Appointment) (*model.Appointment, error)
	UpdateAppointment(context.Context, model.Appointment) (*model.Appointment, error)
	// DeleteAppointment(context.Context, string) error
}

type AppointmentMemotyI interface {
	QuerieMemory

	ExecerMemory
}

type QuerieMemory interface {
}

type ExecerMemory interface {
}
