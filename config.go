package main

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

const (
	prefix = ""
)

type Config struct {
	MongoURI      string `env:"MONGO_URI"`
	MongoPassword string `env:"MONGO_PASSWORD"`
	MongoUser     string `env:"MONGO_USER"`
	MongoDatabase string `env:"MONGO_DATABASE"`
}

func NewConfig(ctx context.Context) (*Config, error) {
	conf := &Config{}

	pl := envconfig.PrefixLookuper(prefix, envconfig.OsLookuper())
	if err := envconfig.ProcessWith(ctx, conf, pl); err != nil {
		return nil, err
	}
	return conf, nil
}
