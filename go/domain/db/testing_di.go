package db

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/sarulabs/di"
	"github.com/teawithsand/webpage/domain/common/dikey"
	"github.com/teawithsand/webpage/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO(teawithsand): here add utils for mock databases for testing

type TestingDB struct {
	DB *TypedDB
}

func (tdb *TestingDB) Close() (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	err = tdb.DB.Database.Drop(ctx)
	return
}

func MakeTestingDB(ctx context.Context, cfg Config) (newdb *TestingDB, err error) {
	v := rand.Int63()
	name := fmt.Sprintf("ftwapi-tdb-%d", v)

	var client *mongo.Client
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURL))
	if err != nil {
		return
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return
	}

	rdb := client.Database(name)

	tdb, err := MakeTypedDB(ctx, rdb)
	if err != nil {
		return
	}

	err = SetupDB(ctx, tdb)
	if err != nil {
		return
	}

	newdb = &TestingDB{DB: tdb}
	return
}

func RegisterTestInDI(builder util.Builder) (err error) {
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
		Name: dikey.DBTestDatabaseDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			rcfg, err := ctn.SafeGet(dikey.DBConfigDI)
			if err != nil {
				return
			}
			tcfg := rcfg.(*Config)

			tdb, err := MakeTestingDB(context.TODO(), *tcfg)
			if err != nil {
				return
			}

			res = tdb
			return
		},
		Close: func(obj interface{}) (err error) {
			err = obj.(*TestingDB).Close()
			return
		},
	})

	builder.Add(util.Def{
		Name: dikey.DBDatabaseDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			var tdb *TestingDB
			err = ctn.Fill(dikey.DBTestDatabaseDI, &tdb)
			if err != nil {
				return
			}

			res = tdb.DB
			return
		},
	})

	return
}
