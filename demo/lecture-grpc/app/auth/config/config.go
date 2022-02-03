package config

type AppConfig struct {
	Host string
	Port int
}

func Load() *AppConfig {
	return &AppConfig{
		Host: "",
		Port: 9000,
	}
}
