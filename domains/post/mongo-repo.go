package post

import (
	"instagram2/db"
)

type MongoPostRepo struct {
}

func NewMongoPostRepo(client *db.Mongo) *MongoPostRepo {
	return &MongoPostRepo{}
}

func (r *MongoPostRepo) GetPosts() []string {
	return []string{"", ""}
}
