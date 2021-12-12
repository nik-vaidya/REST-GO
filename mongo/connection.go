package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func CreateConnection(ctx context.Context, url string) (*mongo.Client, error) {
	options := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func ConnectToDataBase(ctx context.Context, url string) (*mongo.Client, error) {

	client, err := CreateConnection(ctx, url)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
