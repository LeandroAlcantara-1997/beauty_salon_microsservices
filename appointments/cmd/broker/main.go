package main

import (
	"context"
	"log"
	"time"

	broker "github.com/LeandroAlcantara-1997/appointment/internal/api"
	"github.com/LeandroAlcantara-1997/appointment/internal/config"
	"github.com/LeandroAlcantara-1997/appointment/internal/container"
	"github.com/facily-tech/go-core/types"
)

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, types.ContextKey(types.Version), config.NewVersion())
	ctx = context.WithValue(ctx, types.ContextKey(types.StartedAt), time.Now())
	_, dep, err := container.New(ctx)
	if err != nil {
		log.Fatal(err) // log might not be started and because of that dep might not exist
	}

	conn := dep.Components.RabbitMQ
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// Make a channel to receive messages into infinite loop.

	if err := broker.Broker(ch, dep); err != nil {
		log.Println(err)
	}
}
