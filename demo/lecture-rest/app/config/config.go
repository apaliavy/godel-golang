package config

import (
	"time"
)

type DBConf struct {
	Username string
	Password string
	Host     string
	Port     int
	Name     string
	Options  string
}

type APIConf struct {
	Host            string
	Port            int
	GracefulTimeout time.Duration
}

type AppConfig struct {
	DB  DBConf
	API APIConf
}

func Load() *AppConfig {
	return &AppConfig{
		DB: DBConf{
			Username: "postgres",
			Password: "mysecretpassword",
			Host:     "localhost",
			Port:     5432,
			Name:     "users",
			Options:  "sslmode=disable",
		},
		API: APIConf{
			Host:            "localhost",
			Port:            8080,
			GracefulTimeout: 10,
		},
	}
}
