package model

import (
	"time"
)

type Appointment struct {
	ID              string    `bson:"_id,omitempty"`
	UserID          int       `bson:"user_id"`
	SalonID         int       `bson:"salon_id"`
	AppointmentDate time.Time `bson:"appointment_date"`
}

func NewAppointment(appointment UpsertAppointment) Appointment {
	return Appointment(appointment)
}
