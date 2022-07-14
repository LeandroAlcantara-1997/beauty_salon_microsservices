package main

import (
	"context"
	"log"
	"time"

	_ "github.com/LeandroAlcantara-1997/appointment/docs"
	"github.com/LeandroAlcantara-1997/appointment/internal/api"
	"github.com/LeandroAlcantara-1997/appointment/internal/config"
	"github.com/LeandroAlcantara-1997/appointment/internal/container"

	"github.com/facily-tech/go-core/env"
	apiServer "github.com/facily-tech/go-core/http/server"
	"github.com/facily-tech/go-core/types"

	_ "github.com/golang/mock/mockgen/model"
)

// @title           Appointment API
// @version         1.0
// @description     This is a service for make appointments .
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://github.com/LeandroAlcantara-1997
// @contact.email  leandro1997silva97@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1/appointment

func main() {
	// root context of application
	ctx := context.Background()

	ctx = context.WithValue(ctx, types.ContextKey(types.Version), config.NewVersion())
	ctx = context.WithValue(ctx, types.ContextKey(types.StartedAt), time.Now())

	ctx, dep, err := container.New(ctx)
	if err != nil {
		log.Fatal(err) // log might not be started and because of that dep might not exist
	}

	apiConfig := apiServer.Config{}
	err = env.LoadEnv(ctx, &apiConfig, apiServer.PrefixConfig)
	if err != nil {
		log.Fatal(err)
	}

	apiServer.Run(
		ctx,
		apiConfig,
		api.Handler(ctx, dep),
		dep.Components.Log,
	)

	dep.Components.Tracer.Close()
}
