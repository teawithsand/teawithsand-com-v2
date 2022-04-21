package dbutil

import (
	"github.com/teawithsand/webpage/domain/db"
	"go.mongodb.org/mongo-driver/bson"
)

func NewIDQuery(id db.ID) interface{} {
	return struct {
		ID db.ID `bson:"_id"`
	}{
		ID: id,
	}
}

func NewNotQuery(query interface{}) interface{} {
	return bson.D{
		{"$not", query},
	}
}

func NewEQQuery(field string, value interface{}) interface{} {
	return bson.D{
		{field, bson.D{{"$eq", value}}},
	}
}

func NewNEQuery(field string, value interface{}) interface{} {
	return bson.D{
		{field, bson.D{{"$ne", value}}},
	}
}

func NewExistsQuery(field string, exists bool) interface{} {
	return bson.D{
		{field, bson.D{{"$exists", exists}}},
	}
}
