package db

import (
	"context"
	"time"

	"github.com/sarulabs/di"
	"github.com/teawithsand/webpage/domain/common/dikey"
	"github.com/teawithsand/webpage/util"
	"github.com/teawithsand/webpage/util/env"
)

func RegisterInDI(builder util.Builder) (err error) {
	var config Config

	err = util.ReadConfig(&config)
	if err != nil {
		return
	}

	builder.Add(util.Def{
		Name: dikey.DBConfigDI,
		Build: func(ctn util.DI) (interface{}, error) {
			return &config, nil
		},
	})

	builder.Add(util.Def{
		Name: dikey.DBDatabaseDI,

		Build: func(ctn di.Container) (res interface{}, err error) {
			var env *env.Config
			err = ctn.Fill(dikey.EnvDI, &env)
			if err != nil {
				return
			}

			var dbConfig *Config
			err = ctn.Fill(dikey.DBConfigDI, &dbConfig)
			if err != nil {
				return
			}

			tdb, err := MakeDB(context.TODO(), *dbConfig)
			if err != nil {
				return
			}

			res = tdb
			return
		},

		Close: func(obj interface{}) (err error) {
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, time.Second*15)
			defer cancel()

			err = obj.(*TypedDB).Client.Disconnect(ctx)
			return
		},
	})

	return
}
