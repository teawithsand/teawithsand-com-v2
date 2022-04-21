package user

import (
	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/util/dbutil"
)

type RegistrationRepository = dbutil.Repository

func NewRegistrationRepository(collection db.Collection) RegistrationRepository {
	return &dbutil.RepositoryImpl{
		Collection: collection,
		CreateDataChecker: func(data interface{}) (err error) {
			_, ok := data.(Registration)
			if !ok {
				err = dbutil.ErrInvalidData
				return
			}
			return
		},
	}
}
