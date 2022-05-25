package container

import (
	"context"
	"fmt"

	"github.com/LeandroAlcantara-1997/appointment/internal/config"
	mongoConfig "github.com/LeandroAlcantara-1997/appointment/pkg/core/mongo"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/repository"
	app "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/service"
	"github.com/facily-tech/go-core/env"
	"github.com/facily-tech/go-core/log"
	"github.com/facily-tech/go-core/telemetry"
	"github.com/facily-tech/go-core/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type envs struct {
	Mongo mongoConfig.Config
}

// Components are a like service, but it doesn't include business case
// Or domains, but likely used by multiple domains
type components struct {
	Log         log.Logger
	Tracer      telemetry.Tracer
	MongoClient *mongo.Client
	// Include your new components bellow
}

// Services hold the business case, and make the bridge between
// Controllers and Domains
type Services struct {
	Appointments app.ServiceI
}

type Dependency struct {
	Components components
	Services   Services
}

func New(ctx context.Context) (context.Context, *Dependency, error) {
	envs, err := loadEnvs(ctx)
	if err != nil {
		return nil, nil, err
	}

	cmp, err := setupComponents(ctx, envs)
	if err != nil {
		return nil, nil, err
	}

	apService, err := app.NewService(
		cmp.Log,
		repository.NewMongoRepostory(
			cmp.MongoClient,
			envs.Mongo.Database,
			envs.Mongo.Collection,
		),
	)
	if err != nil {
		return nil, nil, err
	}

	srv := Services{
		// include services initialized above here
		apService,
	}

	dep := Dependency{
		Components: *cmp,
		Services:   srv,
	}

	return ctx, &dep, err
}
func loadEnvs(ctx context.Context) (envs, error) {
	mongoDB := mongoConfig.Config{}
	if err := env.LoadEnv(ctx, &mongoDB, mongoConfig.ConfigPrefix); err != nil {
		return envs{}, err
	}

	return envs{
		Mongo: mongoDB,
	}, nil
}

func setupComponents(ctx context.Context, envs envs) (*components, error) {
	version, ok := ctx.Value(types.ContextKey(types.Version)).(*config.Version)
	if !ok {
		return nil, config.ErrVersionTypeAssertion
	}

	telemetryConfig := telemetry.DataDogConfig{
		Version: version.GitCommitHash,
	}

	err := env.LoadEnv(ctx, &telemetryConfig, telemetry.DataDogConfigPrefix)
	if err != nil {
		return nil, err
	}

	tracer, err := telemetry.NewDataDog(telemetryConfig)

	if err != nil {
		return nil, err
	}

	l, err := log.NewLoggerZap(log.ZapConfig{
		Version:           version.GitCommitHash,
		DisableStackTrace: true,
		Tracer:            tracer,
	})

	if err != nil {
		return nil, err
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:27017",
		envs.Mongo.User,
		envs.Mongo.Password,
		envs.Mongo.Host)))

	if err != nil {
		return nil, err
	}
	return &components{
		Log:         l,
		Tracer:      tracer,
		MongoClient: client,
		// include components initialized bellow here
	}, nil
}
