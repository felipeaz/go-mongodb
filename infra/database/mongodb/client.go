package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	C *mongo.Client
}

func Connect(clusterUri string) (*Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(clusterUri))
	if err != nil {
		return nil, err
	}
	return &Client{C: client}, nil
}

func (c Client) Disconnect() error {
	return c.C.Disconnect(context.TODO())
}
