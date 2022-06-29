package transport

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments"
	appErr "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/error"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/model"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/service"
	"github.com/go-kit/kit/transport/amqp"
	"github.com/pkg/errors"
	delivery "github.com/streadway/amqp"
)

const queue = 2

func NewBroke(svc service.AppointmentService, ch amqp.Channel) error {
	wg := new(sync.WaitGroup)
	wg.Add(queue)
	options := []amqp.SubscriberOption{
		amqp.SubscriberErrorEncoder(errorSubscriber),
	}

	createApp := amqp.NewSubscriber(
		appointments.CreateAppointment(svc),
		decodeCreateApp,
		encodeResponseFunc,
		options...,
	).ServeDelivery(ch)

	createMessage, err := ch.Consume("create", "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	makeApp := amqp.NewSubscriber(
		appointments.MakeAppointmentByUser(svc),
		decodeMakeAppointment,
		encodeResponseFunc,
		options...,
	).ServeDelivery(ch)

	makeMessage, err := ch.Consume("make", "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go createchannel(createApp, createMessage, wg)
	go makechannel(makeApp, makeMessage, wg)
	wg.Wait()
	return nil
}

func createchannel(del func(del *delivery.Delivery), message <-chan delivery.Delivery, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range message {
		del(&d)
		log.Printf("Received a message: %s", d.Body)
	}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}

func makechannel(del func(del *delivery.Delivery), message <-chan delivery.Delivery, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range message {
		del(&d)
		log.Printf("Received a message: %s", d.Body)
	}
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

func decodeMakeAppointment(_ context.Context, r *delivery.Delivery) (interface{}, error) {
	var app model.MakeAppointment

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
