package post

import (
	"context"
	"instagram2/db"
)

type MongoPostRepo struct {
	Mongo *db.Mongo
}

func NewMongoPostRepo(mongo *db.Mongo) *MongoPostRepo {
	return &MongoPostRepo{
		Mongo: mongo,
	}
}

func (r *MongoPostRepo) InsertPost(post Post) string {
	r.Mongo.Client.InsertOne(context.TODO(), post)
	return "creado con exito"
}
