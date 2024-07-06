package post

type Post struct {
	UserId      string   `json:"user_id" bson:"user_id"`
	Description string   `json:"description" bson:"description"`
	Likes       []string `json:"likes" bson:"likes"`
	Comments    []string `json:"comments" bson:"comments"`
}
