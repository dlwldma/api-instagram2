package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
}

func NewMongoDb() *Mongo {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("uri"))
	return &Mongo{
		Client: client,
	}
}
