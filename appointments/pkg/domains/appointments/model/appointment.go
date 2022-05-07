package model

import (
	"time"
)

type Appointment struct {
	ID              string
	UserID          int
	SalonID         int
	AppointmentDate time.Time
}

func NewAppointment(appointment UpsertAppointment) Appointment {
	return Appointment{
		UserID:          appointment.UserID,
		SalonID:         appointment.SalonID,
		AppointmentDate: appointment.AppointmentDate,
	}
}
