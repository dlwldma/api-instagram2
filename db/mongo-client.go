package db

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Collection
}

func NewMongoDb(collection string) *Mongo {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	uri := os.Getenv("MONGO_URI")
	db_name := os.Getenv("MONGO_DB_NAME")
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	return &Mongo{
		Client: client.Database(db_name).Collection(collection),
	}
}
