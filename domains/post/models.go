package post

import "time"

type Post struct {
	Id          string    `json:"_id" bson:"_id"`
	UserId      string    `json:"user_id" bson:"user_id"`
	Date        time.Time `json:"date" bson:"date"`
	Description string    `json:"description" bson:"description"`
	Image_urls  []string  `json:"image_url" bson:"image_url"`
	Likes       []string  `json:"likes" bson:"likes"`
	Comments    []string  `json:"comments" bson:"comments"`
}
