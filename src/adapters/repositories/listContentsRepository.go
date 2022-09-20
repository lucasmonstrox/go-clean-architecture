package repositories

import (
	"app/src/domain/entities"
	mongoGoDriver "app/src/ports/mongo-go-driver/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupListContentsRepository(client *mongo.Client) func() []entities.Content {
	return func() []entities.Content {
		collection := client.Database("spalla").Collection("contents")
		contents := make([]entities.Content, 0)
		cursor, _ := collection.Find(context.TODO(), bson.D{})
		var contentsDocuments []mongoGoDriver.Content
		cursor.All(context.TODO(), &contentsDocuments)
		for _, contentDocument := range contentsDocuments {
			content := entities.Content{
				Id:          contentDocument.Id,
				Title:       contentDocument.Title,
				Description: contentDocument.Description,
				Thumbnail:   contentDocument.Thumbnail,
			}
			contents = append(contents, content)
		}
		return contents
	}
}
