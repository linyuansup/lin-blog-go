package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Dao struct {
	collection *mongo.Collection
}

func (d *Dao) findOne(filter map[string]interface{}) *mongo.SingleResult {
	var f bson.D
	for k, v := range filter {
		f = append(f, bson.E{Key: k, Value: v})
	}
	return d.collection.FindOne(context.Background(), f)
}

func (d *Dao) findMany(filter map[string]interface{}, limit, page int64) (*mongo.Cursor, error) {
	var f bson.D
	for k, v := range filter {
		f = append(f, bson.E{Key: k, Value: v})
	}
	findOptions := &options.FindOptions{}
	if limit > 0 {
		findOptions.SetLimit(limit)
		findOptions.SetSkip((limit * page) - page)
	}
	cur, err := d.collection.Find(context.Background(), f, findOptions)
	if err == nil {
		err = cur.Err()
	}
	return cur, err
}
