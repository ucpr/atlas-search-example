package main

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

const (
	prefix = "MONGO_"
)

type Config struct {
	MongoURI      string `env:"URI"`
	MongoPassword string `env:"PASSWORD"`
	MongoUser     string `env:"USER"`
	MongoDatabase string `env:"DATABASE"`
}

func NewConfig(ctx context.Context) (*Config, error) {
	conf := &Config{}

	pl := envconfig.PrefixLookuper(prefix, envconfig.OsLookuper())
	if err := envconfig.ProcessWith(ctx, conf, pl); err != nil {
		return nil, err
	}
	return conf, nil
}
