package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blog struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title"`
	Summary  string             `bson:"summary"`
	Content  string             `bson:"content"`
	Time     string             `bson:"time"`
	Category primitive.ObjectID `bson:"category"`
	Picture  string             `bson:"picture"`
}
