package mongogodriver

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	clientOpts := options.Client().SetHosts(
		[]string{"localhost:27017"},
	).SetAuth(
		options.Credential{
			AuthSource:    "spalla",
			AuthMechanism: "SCRAM-SHA-256",
			Username:      "spalla",
			Password:      "spalla",
		},
	)
	client, _ := mongo.Connect(context.TODO(), clientOpts)
	return client
}
