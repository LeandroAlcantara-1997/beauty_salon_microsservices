package redis

const ConfigPrefix = "REDIS_"

type Config struct {
	Host     string `env:"HOST, required"`
	Password string `env:"PASSWORD, required"`
}
