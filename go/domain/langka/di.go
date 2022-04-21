package langka

import (
	"github.com/sarulabs/di"
	"github.com/teawithsand/webpage/domain/common/dikey"
	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/util"
)

func RegisterInDI(builder util.Builder) (err error) {
	builder.Add(di.Def{
		Name: dikey.LangkaWordSetRepositoryDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			var db *db.TypedDB
			err = ctn.Fill(dikey.DBDatabaseDI, &db)
			if err != nil {
				return
			}

			res = NewWordSetRepository(db.WordSets)
			return
		},
	})

	builder.Add(di.Def{
		Name: dikey.LangkaWordSetServiceDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			var svc WordSetService

			err = ctn.Fill(dikey.LangkaWordSetRepositoryDI, &svc.WordSetRepository)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.CheckerDI, &svc.Checker)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.SluggifierDI, &svc.Sluggifier)
			if err != nil {
				return
			}

			// validators are not managed by DI for now
			svc.CreateWordSetDataValidator = DefaultValidators.WordSetCreateDataValidator

			res = &svc
			return
		},
	})

	return
}
