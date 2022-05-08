package mongo

const ConfigPrefix = "MONGO_"

type Config struct {
	Host       string `env:"HOST, required"`
	User       string `env:"USER, required"`
	Password   string `env:"PASSWORD, required"`
	Collection string `env:"COLLECTION, required"`
}
