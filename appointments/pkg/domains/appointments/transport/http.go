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
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

var validate = validator.New()

func NewHTTPHandler(svc service.AppointmentServiceI) stdHTTP.Handler {
	options := []http.ServerOption{
		http.ServerErrorEncoder(errorHandler),
	}

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

	availableApp := http.NewServer(
		appointments.AvailableAppointment(svc),
		decodeAvailableApp,
		codeHTTP{200}.encodeResponse,
		options...,
	)

	deleteApp := http.NewServer(
		appointments.DeleteAppointment(svc),
		decodeDeleteApp,
		codeHTTP{204}.encodeResponse,
		options...,
	)

	cancelApp := http.NewServer(
		appointments.CancelAppointment(svc),
		decodeCancelApp,
		codeHTTP{204}.encodeResponse,
		options...,
	)

	r := chi.NewRouter()

	r.Get("/{id}", findAppByID.ServeHTTP)
	r.Get("/", findAllApp.ServeHTTP)
	r.Get("/user/{id}", findAppByUserID.ServeHTTP)
	r.Get("/salon/{id}", findAppBySalonID.ServeHTTP)
	r.Get("/available", availableApp.ServeHTTP)
	r.Put("/{id}", updateApp.ServeHTTP)
	r.Put("/{id}/{user}", cancelApp.ServeHTTP)
	r.Delete("/{id}", deleteApp.ServeHTTP)

	return r
}

func decodeFindAppByID(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var app model.FindAppointmentsByIDRequest
	if app.ID = chi.URLParam(r, "id"); app.ID == "" {
		return nil, appErr.ErrInvalidPath
	}
	return app, nil
}

// ShowAccount godoc
// @Summary      Update an appointment
// @Description  get string by ID and body for update
// @Tags         appointment
// @Accept       json
// @Produce      json
// @Failure      500  {}  Sorry, we cannot create a new appointment
// @Failure      404  {}  Appointment not found
// @Failure      400  {}  An error happened in database
// @Failure      400  {}  Cannot read path Invalid body
// @Success      200  {object}   model.AppResponse
// @Param        id   path      string  true  "Appointment ID"
// @Param appointment body string true "body appointment"
// SchemaExample({\n"user_id": 1,\n"salon_id": 2,\n"appointment_date": "2022-06-23T21:12:02.000000001Z"\n})
// @Router       /appointment/{id} [put]
func decodeUpdateApp(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var app model.UpsertAppointment
	if app.ID = chi.URLParam(r, "id"); app.ID == "" {
		return nil, appErr.ErrInvalidPath
	}

	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		return nil, appErr.ErrInvalidBody
	}

	if err := validate.Struct(app); err != nil {
		return nil, errors.Wrap(appErr.ErrInvalidBody, err.Error())
	}

	return app, nil
}

// ShowAccount godoc
// @Summary      Get all appointments
// @Description  get all appointments
// @Tags         appointment
// @Accept       json
// @Produce      json
// @Failure      500  {}  Sorry, we cannot create a new appointment
// @Failure      404  {}  Appointment not found
// @Failure      400  {}  An error happened in database
// @Success      200  {array}   model.AppResponse
// @Router       /appointment [get]
func decodeAllApp(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	return nil, nil
}

// ShowAccount godoc
// @Summary      Get appointments by user ids
// @Description  get string by user ID and return an appointment
// @Tags         appointment
// @Accept       json
// @Produce      json
// @Failure      500  {}  Sorry, we cannot create a new appointment
// @Failure      404  {}  Appointment not found
// @Failure      400  {}  An error happened in database
// @Failure      400  {}  Cannot read path
// @Success      200  {array}   model.AppResponse
// @Param        id   path      int  true  "User ID"
// @Router       /appointment/user/{id} [get]
func decodeAppByUser(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var (
		app model.FindAppByUser
		err error
	)
	if app.ID, err = strconv.Atoi(chi.URLParam(r, "id")); err != nil {
		return nil, appErr.ErrInvalidPath
	}

	return app, nil
}

// ShowAccount godoc
// @Summary      Get appointments by user ids
// @Description  get string by salon ID and return an appointment
// @Tags         appointment
// @Accept       json
// @Produce      json
// @Failure      500  {}  Sorry, we cannot create a new appointment
// @Failure      404  {}  Appointment not found
// @Failure      400  {}  An error happened in database
// @Failure      400  {}  Cannot read path
// @Success      200  {array}   model.AppResponse
// @Param        id   path      int  true  "Salon ID"
// @Router       /appointment/salon/{id} [get]
func decodeAppBySalon(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var (
		app model.FindAppBySalon
		err error
	)
	if app.ID, err = strconv.Atoi(chi.URLParam(r, "id")); err != nil {
		return nil, appErr.ErrInvalidPath
	}

	return app, nil
}

// ShowAccount godoc
// @Summary      Delete appointments by id
// @Description  get string by ID and delete an appointment
// @Tags         appointment
// @Accept       json
// @Produce      json
// @Failure      500  {}  Sorry, we cannot create a new appointment
// @Failure      404  {}  Appointment not found
// @Failure      400  {}  An error happened in database
// @Failure      400  {}  Cannot read path
// @Success      204
// @Param        id   path      string  true  "Appointment ID"
// @Router       /appointment/{id} [delete]
func decodeDeleteApp(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var app model.DeleteAppointment
	if app.ID = chi.URLParam(r, "id"); app.ID == "" {
		return nil, appErr.ErrInvalidPath
	}
	return app, nil
}

// ShowAccount godoc
// @Summary      Get available appointments
// @Description  get all available appointments
// @Tags         appointment
// @Accept       json
// @Produce      json
// @Failure      500  {}  Sorry, we cannot create a new appointment
// @Failure      404  {}  Appointment not found
// @Failure      400  {}  An error happened in database
// @Success      200  {array}   model.AppResponse
// @Router       /appointment/available [get]
func decodeAvailableApp(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	return nil, nil
}

func decodeCancelApp(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var (
		app model.MakeAppointment
		err error
	)
	if app.ID = chi.URLParam(r, "id"); app.ID == "" {
		return nil, appErr.ErrInvalidPath
	}

	if app.UserID, err = strconv.Atoi(chi.URLParam(r, "user")); err != nil {
		return nil, appErr.ErrInvalidPath
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
