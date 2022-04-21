package dbutil

import (
	"go.mongodb.org/mongo-driver/bson"
)

func QueryAnd(queries ...interface{}) interface{} {
	if len(queries) == 0 {
		return nil
	}
	var filtered []interface{}
	for _, q := range queries {
		if q != nil {
			filtered = append(filtered, q)
		}
	}
	return bson.M{
		"$and": bson.A(filtered),
	}
}

func NewORQuery(queries ...interface{}) interface{} {
	if len(queries) == 0 {
		return nil
	}
	var filtered []interface{}
	for _, q := range queries {
		if q != nil {
			filtered = append(filtered, q)
		}
	}
	return bson.M{
		"$or": bson.A(filtered),
	}
}
