package config

import (
	"time"
)

type APIConf struct {
	Host            string
	Port            int
	GracefulTimeout time.Duration
}

type RPCConf struct {
	Host string
	Port int
}

type AppConfig struct {
	API             APIConf
	UsersServiceRPC RPCConf
	AuthServiceRPC  RPCConf
}

func Load() *AppConfig {
	return &AppConfig{
		API: APIConf{
			Host:            "",
			Port:            8080,
			GracefulTimeout: 10,
		},
		UsersServiceRPC: RPCConf{
			Host: "example-users",
			Port: 9001,
		},
		AuthServiceRPC: RPCConf{
			Host: "example-auth",
			Port: 9000,
		},
	}
}
