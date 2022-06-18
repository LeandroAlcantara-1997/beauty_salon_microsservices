package container

import (
	"context"
	"fmt"

	"github.com/LeandroAlcantara-1997/appointment/internal/config"
	mongoConfig "github.com/LeandroAlcantara-1997/appointment/pkg/core/mongo"
	rabbitConfig "github.com/LeandroAlcantara-1997/appointment/pkg/core/rabbitmq"
	redisConfig "github.com/LeandroAlcantara-1997/appointment/pkg/core/redis"
	"github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/repository"
	app "github.com/LeandroAlcantara-1997/appointment/pkg/domains/appointments/service"
	"github.com/facily-tech/go-core/env"
	"github.com/facily-tech/go-core/log"
	"github.com/facily-tech/go-core/telemetry"
	"github.com/facily-tech/go-core/types"
	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type envs struct {
	Mongo  mongoConfig.Config
	Redis  redisConfig.Config
	Rabbit rabbitConfig.Config
}

// Components are a like service, but it doesn't include business case
// Or domains, but likely used by multiple domains
type components struct {
	Log         log.Logger
	Tracer      telemetry.Tracer
	MongoClient *mongo.Client
	RedisClient *redis.Client
	RabbitMQ    *amqp.Connection
	// Include your new components bellow
}

// Services hold the business case, and make the bridge between
// Controllers and Domains
type Services struct {
	Appointments app.AppointmentService
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
		repository.NewRedisRepository(
			cmp.RedisClient,
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

	redisDB := redisConfig.Config{}
	if err := env.LoadEnv(ctx, &redisDB, redisConfig.ConfigPrefix); err != nil {
		return envs{}, err
	}
	rabbit := rabbitConfig.Config{}
	if err := env.LoadEnv(ctx, &rabbit, rabbitConfig.ConfigPrefix); err != nil {
		return envs{}, err
	}

	return envs{
		Mongo:  mongoDB,
		Redis:  redisDB,
		Rabbit: rabbit,
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

	clientMongo, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:27017",
			envs.Mongo.User,
			envs.Mongo.Password,
			envs.Mongo.Host)))

	if err != nil {
		return nil, err
	}

	clientRedis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", envs.Redis.Host),
		Password: envs.Redis.Password,
	})

	clientRabbitMQ, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@message-broker",
		envs.Rabbit.User, envs.Rabbit.Password))
	if err != nil {
		return nil, err
	}

	return &components{
		Log:         l,
		Tracer:      tracer,
		MongoClient: clientMongo,
		RedisClient: clientRedis,
		RabbitMQ:    clientRabbitMQ,
		// include components initialized bellow here
	}, nil
}
