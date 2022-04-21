package dbutil

import (
	"context"
	"fmt"

	"github.com/teawithsand/webpage/domain/db"
	"go.mongodb.org/mongo-driver/bson"
)

func DebugCollection(ctx context.Context, coll db.Collection) (err error) {
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return
	}

	res := []interface{}{}
	err = cursor.All(ctx, &res)
	if err != nil {
		return
	}

	fmt.Printf("Got %d results:\n%+#v\n", len(res), res)
	return
}

func DebugCollectionFirst(ctx context.Context, coll db.Collection, deserializeTarget interface{}) (err error) {
	err = coll.FindOne(ctx, bson.D{}).Decode(deserializeTarget)
	if err != nil {
		return
	}

	fmt.Printf("Got:\n%+#v\n", deserializeTarget)
	return
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
