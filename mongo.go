package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	cli *mongo.Client
	db  string
}

func NewClient(ctx context.Context, cfg *Config) (*Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(cfg.MongoURI).
		SetAuth(options.Credential{
			Username: cfg.MongoUser,
			Password: cfg.MongoPassword,
		}).
		SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping your deployment: %w", err)
	}
	log.Println("successfully connected to MongoDB")

	return &Client{
		cli: client,
		db:  cfg.MongoDatabase,
	}, nil
}

func (c *Client) Disconnect(ctx context.Context) error {
	return c.cli.Disconnect(ctx)
}

func (c *Client) collection(col string) *mongo.Collection {
	return c.cli.Database(c.db).Collection(col)
}

func (c *Client) Movies() *mongo.Collection {
	return c.collection("movies")
}
