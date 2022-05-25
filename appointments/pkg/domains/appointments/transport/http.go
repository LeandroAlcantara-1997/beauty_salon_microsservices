package transport

import (
	"context"
	"encoding/json"
	"log"
	stdHTTP "net/http"
	"strconv"

	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments"
	appErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(svc service.ServiceI) stdHTTP.Handler {
	options := []http.ServerOption{
		http.ServerErrorEncoder(errorHandler),
	}

	createApp := http.NewServer(
		appointments.CreateAppointmentByID(svc),
		decodeCreateApp,
		codeHTTP{201}.encodeResponse,
		options...,
	)

	updateApp := http.NewServer(
		appointments.UpdateAppointmentByUser(svc),
		decodeUpdateApp,
		codeHTTP{200}.encodeResponse,
		options...,
	)

	findAppByID := http.NewServer(
		appointments.FindAppointmentByID(svc),
		decodeFindAppByID,
		codeHTTP{200}.encodeResponse,
		options...,
	)

	findAllApp := http.NewServer(
		appointments.FindAllAppointment(svc),
		decodeAllApp,
		codeHTTP{200}.encodeResponse,
		options...,
	)

	findAppByUserID := http.NewServer(
		appointments.FindAppointmentByUser(svc),
		decodeAppByUser,
		codeHTTP{200}.encodeResponse,
		options...,
	)

	findAppBySalonID := http.NewServer(
		appointments.FindAppointmentBySalon(svc),
		decodeAppBySalon,
		codeHTTP{200}.encodeResponse,
		options...,
	)

	r := chi.NewRouter()

	r.Post("/", createApp.ServeHTTP)
	r.Get("/{id}", findAppByID.ServeHTTP)
	r.Get("/", findAllApp.ServeHTTP)
	r.Get("/user/{id}", findAppByUserID.ServeHTTP)
	r.Get("/salon/{id}", findAppBySalonID.ServeHTTP)
	r.Put("/{id}", updateApp.ServeHTTP)

	return r
}

func decodeFindAppByID(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var app model.FindAppointmentsByIDRequest
	if app.ID = chi.URLParam(r, "id"); app.ID == "" {
		return nil, appErr.ErrInvalidPath
	}
	return app, nil
}

func decodeCreateApp(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var app model.UpsertAppointment
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		return nil, appErr.ErrInvalidBody
	}

	return app, nil
}

func decodeUpdateApp(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var app model.UpsertAppointment
	if app.ID = chi.URLParam(r, "id"); app.ID == "" {
		return nil, appErr.ErrInvalidPath
	}

	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		return nil, appErr.ErrInvalidBody
	}

	return app, nil
}

func decodeAllApp(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	return nil, nil
}

func decodeAppByUser(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var (
		app model.FindAppByUser
		err error
	)
	if app.ID, err = strconv.Atoi(chi.URLParam(r, "id")); err != nil {
		return nil, err
	}

	return app, nil
}

func decodeAppBySalon(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var (
		app model.FindAppBySalon
		err error
	)
	if app.ID, err = strconv.Atoi(chi.URLParam(r, "id")); err != nil {
		return nil, err
	}

	return app, nil
}

type codeHTTP struct {
	int
}

func (c codeHTTP) encodeResponse(_ context.Context, w stdHTTP.ResponseWriter, input interface{}) error {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(c.int)
	return json.NewEncoder(w).Encode(input)
}

func errorHandler(_ context.Context, err error, w stdHTTP.ResponseWriter) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	resp, code := appErr.RESTErrorBussines.ErrorProcess(err)

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": resp}); err != nil {
		log.Printf("Encoding error, nothing much we can do: %v", err)
	}
}
