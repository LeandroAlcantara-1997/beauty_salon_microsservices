package appointments

import (
	"context"

	appErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

func CreateAppointmentByID(svc service.ServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(model.UpsertAppointment)
		if !ok {
			return nil, errors.Wrap(appErr.ErrTypeAssertion, "cannot convert request -> UpsertAppointment")
		}

		appResponse, err := svc.CreateAppointment(ctx, req)
		if err != nil {
			return nil, err
		}

		return appResponse, nil
	}
}

func FindAppointmentByID(svc service.ServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(model.FindAppointmentsByIDRequest)
		if !ok {
			return nil, errors.Wrap(appErr.ErrTypeAssertion, "cannot convert request -> FindAppointmentsByIDRequest")
		}

		appResponse, err := svc.FindAppByID(ctx, req)
		if err != nil {
			return nil, err
		}

		return appResponse, nil
	}
}

func FindAllAppointment(svc service.ServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		appResponse, err := svc.FindAllAppointments(ctx)
		if err != nil {
			return nil, err
		}

		return appResponse, nil
	}
}

func FindAppointmentByUser(svc service.ServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(model.FindAppByUser)
		if !ok {
			return nil, errors.Wrap(appErr.ErrTypeAssertion, "cannot convert request -> FindAppByUser")
		}

		appResponse, err := svc.FindAppByUserID(ctx, req)
		if err != nil {
			return nil, err
		}

		return appResponse, nil
	}
}

func FindAppointmentBySalon(svc service.ServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(model.FindAppBySalon)
		if !ok {
			return nil, errors.Wrap(appErr.ErrTypeAssertion, "cannot convert request -> FindAppBySalon")
		}

		appResponse, err := svc.FindAppBySalonID(ctx, req)
		if err != nil {
			return nil, err
		}

		return appResponse, nil
	}
}

func UpdateAppointmentByUser(svc service.ServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(model.UpsertAppointment)
		if !ok {
			return nil, errors.Wrap(appErr.ErrTypeAssertion, "cannot convert request -> UpsertAppointment")
		}

		appResponse, err := svc.UpdateAppointment(ctx, req)
		if err != nil {
			return nil, err
		}

		return appResponse, nil
	}
}
