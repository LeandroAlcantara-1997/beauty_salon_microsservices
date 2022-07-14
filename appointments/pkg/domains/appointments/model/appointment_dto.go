package model

import (
	"time"
)

type UpsertAppointment struct {
	ID              string    `json:"id,omitempty" example:"62b65300e1d7eab1ea9a681d"`
	UserID          int       `json:"user_id" example:"1"`
	SalonID         int       `json:"salon_id" validate:"required" example:"1"`
	AppointmentDate time.Time `json:"appointment_date" validate:"required" example:"2022-06-23T21:12:02.000000001Z"`
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
	ID              string    `json:"id" example:"62b65300e1d7eab1ea9a681d"`
	UserID          int       `json:"user_id" example:"1"`
	SalonID         int       `json:"salon_id" example:"1"`
	AppointmentDate time.Time `json:"appointment_date" example:"2022-06-23T21:12:02.000000001Z"`
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
