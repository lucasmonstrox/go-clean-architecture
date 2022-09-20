package mongoGoDriver

type Content struct {
	Id          string `bson:"_id,omitempty"`
	Title       string `bson:"title,omitempty"`
	Description string `bson:"description,omitempty"`
	Thumbnail   string `bson:"thumbnail,omitempty"`
}
