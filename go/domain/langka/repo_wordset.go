package langka

import (
	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/util/dbutil"
)

type WordSetRepository = dbutil.Repository

func NewWordSetRepository(collection db.Collection) WordSetRepository {
	return &dbutil.RepositoryImpl{
		Collection: collection,
		CreateDataChecker: func(data interface{}) (err error) {
			_, ok := data.(WordSet)
			if !ok {
				err = dbutil.ErrInvalidData
			}
			return
		},
	}
}
