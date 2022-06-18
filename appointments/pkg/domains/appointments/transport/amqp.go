package transport

import (
	"context"
	"encoding/json"
	"log"

	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments"
	appErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/service"
	"github.com/go-kit/kit/transport/amqp"
	"github.com/pkg/errors"
	delivery "github.com/streadway/amqp"
)

func NewBroke(svc service.AppointmentService, ch amqp.Channel) error {
	options := []amqp.SubscriberOption{
		amqp.SubscriberErrorEncoder(errorSubscriber),
	}

	createApp := amqp.NewSubscriber(
		appointments.CreateAppointment(svc),
		decodeCreateApp,
		encodeResponseFunc,
		options...,
	)

	createDelivery := createApp.ServeDelivery(ch)
	me, err := ch.Consume("create", "", false, false, false, false, nil)
	if err != nil {
		return nil
	}

	channel(createDelivery, me)
	return nil
}

func channel(del func(del *delivery.Delivery), message <-chan delivery.Delivery) {
	forever := make(chan bool)
	go func() {
		for me := range message {
			// For example, show received message in a console.
			del(&me)
		}
	}()
	<-forever
}

func decodeCreateApp(_ context.Context, r *delivery.Delivery) (interface{}, error) {
	var app model.UpsertAppointment
	if err := json.Unmarshal(r.Body, &app); err != nil {
		return nil, appErr.ErrInvalidBody
	}
	if err := validate.Struct(app); err != nil {
		return nil, errors.Wrap(appErr.ErrInvalidBody, err.Error())
	}
	return app, nil
}

func encodeResponseFunc(ctx context.Context, p *delivery.Publishing, input interface{}) error {
	var err error
	p.Body, err = json.Marshal(input)
	if err != nil {
		return err
	}
	return nil
}

func errorSubscriber(_ context.Context, err error, deliv *delivery.Delivery, ch amqp.Channel, p *delivery.Publishing) {
	resp, _ := appErr.RESTErrorBussines.ErrorProcess(err)
	p.Body, err = json.Marshal(map[string]string{"error": resp})
	if err != nil {
		log.Printf("Encoding error, nothing much we can do: %v", err)
	}

	if err := deliv.Ack(true); err != nil {
		log.Printf("Cannot be return a response %v", err)
	}
}
