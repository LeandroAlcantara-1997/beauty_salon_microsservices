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
	ID string `json:"id" validate:"required"`
}

type FindAppointmentsByIDRequest struct {
	ID string `json:"id" validate:"required"`
}

type AppResponse struct {
	ID              string    `json:"id"`
	UserID          int       `json:"user_id"`
	SalonID         int       `json:"salon_id"`
	AppointmentDate time.Time `json:"appointment_date"`
}

type MakeAppointment struct {
}

func NewAppResponse(appointment Appointment) AppResponse {
	return AppResponse(appointment)
}
