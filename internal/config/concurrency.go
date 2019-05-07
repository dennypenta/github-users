package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	ConcurrenyLimit int `default:"5"`
}

func New() (Env, error) {
	var env Env
	err := envconfig.Process("", &env)

	return env, err
}
