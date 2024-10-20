package infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client

func InitDB(ctx context.Context, host, port, db string) error {
	clientOptions := options.Client().ApplyURI("").SetMaxPoolSize(100)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	MongoDB = client
	return nil
}
