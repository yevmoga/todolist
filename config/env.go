package config

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

type Configuration struct {
	DatabaseConnection string `env:"DATABASE_CONNECTION,required"`
}

func NewEnv() (config Configuration, err error) {
	err = env.Parse(&config)

	logrus.Info("successful database connection")
	return
}
