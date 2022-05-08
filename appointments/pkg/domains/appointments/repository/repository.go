package repository

import (
	"context"

	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
)

type AppointmentRepositoryI interface {
	Querier

	Execer
}

type Querier interface {
	FindAll(context.Context) ([]model.Appointment, error)
	FindAppointmentByID(context.Context, string) (*model.Appointment, error)
	FindAppointmentByUserID(context.Context, int) ([]model.Appointment, error)
	FindAppointmentBySalonID(context.Context, int) ([]model.Appointment, error)
	MakeAppointment(context.Context) ([]model.Appointment, error)
}

type Execer interface {
	Create(context.Context, model.Appointment) (*model.Appointment, error)
	Update(context.Context, model.Appointment) (*model.Appointment, error)
}
