package api

import (
	"github.com/LeandroAlcantara-1997/appointment/internal/container"
	transport "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/transport"
	"github.com/streadway/amqp"
)

func Broker(ch *amqp.Channel, dep *container.Dependency) error {
	if err := transport.NewBroker(dep.Services.Appointments, ch); err != nil {
		return err
	}

	return nil
}
