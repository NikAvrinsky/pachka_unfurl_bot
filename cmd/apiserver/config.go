package apiserver

type Config struct {
	BindAddr    string
	LogLevel    string
	DatabaseUrl string
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
