package user

import (
	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/util/dbutil"
)

type UserRepository = dbutil.Repository

func NewUserRepository(collection db.Collection) UserRepository {
	return &dbutil.RepositoryImpl{
		Collection: collection,
		CreateDataChecker: func(data interface{}) (err error) {
			_, ok := data.(User)
			if !ok {
				err = dbutil.ErrInvalidData
			}
			return
		},
	}
}
