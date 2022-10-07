package mongoGoDriver

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Done      bool               `bson:"done"`
	CreatedAt time.Time          `bson:"createdAt"`
	DoneAt    *time.Time         `bson:"doneAt"`
}
