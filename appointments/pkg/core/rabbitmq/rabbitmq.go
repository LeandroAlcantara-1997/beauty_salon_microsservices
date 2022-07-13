package rabbitmq

const ConfigPrefix = "RABBIT_"

type Config struct {
	User     string `env:"USER, required"`
	Password string `env:"PASSWORD, required"`
}
