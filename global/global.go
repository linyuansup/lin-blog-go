package global

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	uri                string = "mongodb://127.0.0.1:27017"
	BlogCollection     *mongo.Collection
	CategoryCollection *mongo.Collection
)

func init() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	database := client.Database("blog")
	BlogCollection = database.Collection("blog")
	CategoryCollection = database.Collection("category")
}
