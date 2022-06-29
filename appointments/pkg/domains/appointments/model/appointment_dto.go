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

type FindAppointmentsByIDRequest struct {
	ID string `json:"id"`
}

type FindAppByUser struct {
	ID int `json:"id"`
}

type FindAppBySalon struct {
	ID int `json:"id"`
}

type AppResponse struct {
	ID              string    `json:"id"`
	UserID          int       `json:"user_id"`
	SalonID         int       `json:"salon_id"`
	AppointmentDate time.Time `json:"appointment_date"`
}

type MakeAppointment struct {
	ID     string `json:"id" validate:"required"`
	UserID int    `json:"user_id" validate:"required"`
}

func NewAppResponse(appointment Appointment) AppResponse {
	return AppResponse(appointment)
}

func NewAppResponseSlice(appointment []Appointment) []AppResponse {
	app := make([]AppResponse, 0)
	for _, ap := range appointment {
		app = append(app, NewAppResponse(ap))
	}
	return app
}
