package repositories

import (
	"app/src/domain/entities"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupCreateTodoRepository(client *mongo.Client) func(input entities.CreateTodoInput) entities.Todo {
	collection := client.Database("spalla").Collection("todo")
	return func(input entities.CreateTodoInput) entities.Todo {
		createdAt := time.Now()
		done := false
		res, _ := collection.InsertOne(context.TODO(), bson.D{{"title", input.Title}, {"done", done}, {"createdAt", createdAt}})
		todo := entities.Todo{
			Id:        res.InsertedID.(primitive.ObjectID).String(),
			Title:     input.Title,
			Done:      done,
			CreatedAt: createdAt,
		}
		return todo
	}
}
