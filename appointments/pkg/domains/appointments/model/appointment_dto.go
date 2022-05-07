package model

import (
	"time"
)

type UpsertAppointment struct {
	ID              string    `json:"id,omitempty"`
	UserID          int       `json:"user_id"`
	SalonID         int       `json:"salon_id" validate:"required"`
	AppointmentDate time.Time `json:"appointment_date" validate:"required"`
}

type DeleteAppointment struct {
	ID string `json:"id"`
}

type AppointmentResponse struct {
	ID              string    `json:"id"`
	UserID          int       `json:"user_id"`
	SalonID         int       `json:"salon_id"`
	AppointmentDate time.Time `json:"appointment_date"`
}

func NewAppointmentResponse(appointment Appointment) AppointmentResponse {
	return AppointmentResponse(appointment)
}
