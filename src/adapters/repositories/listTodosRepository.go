package repositories

import (
	"app/src/domain/entities"
	mongoGoDriver "app/src/ports/mongo-go-driver/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupListTodosRepository(client *mongo.Client) func() []entities.Todo {
	collection := client.Database("spalla").Collection("todo")
	return func() []entities.Todo {
		todos := make([]entities.Todo, 0)
		cursor, _ := collection.Find(context.TODO(), bson.D{})
		var todosDocuments []mongoGoDriver.Todo
		cursor.All(context.TODO(), &todosDocuments)
		for _, todoDocument := range todosDocuments {
			todo := entities.Todo{
				Id:        todoDocument.Id,
				Title:     todoDocument.Title,
				Done:      todoDocument.Done,
				CreatedAt: todoDocument.CreatedAt,
				DoneAt:    todoDocument.DoneAt,
			}
			todos = append(todos, todo)
		}
		return todos
	}
}
