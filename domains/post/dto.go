package post

type CreatePostDto struct {
	UserId      string   `json:"user_id" bson:"user_id"`
	Description string   `json:"description" bson:"description"`
	ImageUris64 []string `json:"image_64" bson:"image_64"`
}
