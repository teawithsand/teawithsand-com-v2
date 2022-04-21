package fixture

import (
	"context"

	"github.com/teawithsand/webpage/domain/db"
)

// Applies fixtures to DB given.
func Apply(ctx context.Context, tdb *db.TypedDB) (err error) {
	_, err = tdb.Users.InsertOne(ctx, MockUserOne)
	if err != nil {
		return
	}

	_, err = tdb.WordSets.InsertOne(ctx, MockWordSetOne)
	if err != nil {
		return
	}

	return
}
