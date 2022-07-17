package splunk

const ConfigPrefix = "SPLUNK_"

type Config struct {
	Host       string `env:"HOST, required"`
	Token      string `env:"TOKEN, required"`
	Source     string `env:"SOURCE, required"`
	SourceType string `env:"SOURCETYPE, required"`
	Index      string `env:"INDEX,required"`
	Port       string `env:"PORT, required"`
}
