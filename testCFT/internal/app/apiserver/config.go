package apiserver

type Config struct {
	Bind_addr string `toml:"bind_addr"`
}

func NewConfig() *Config {
	return &Config{}
}
