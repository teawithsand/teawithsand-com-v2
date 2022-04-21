package domain

import (
	"github.com/teawithsand/webpage/domain/common"
	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/domain/langka"
	"github.com/teawithsand/webpage/domain/user"
	"github.com/teawithsand/webpage/domain/webapp"
	"github.com/teawithsand/webpage/util"
	"github.com/teawithsand/webpage/util/env"
)

func ConstructDI() (di util.DI, err error) {
	builder := util.NewDIBuilder()

	envCfg := env.ReadConfig()

	err = common.RegisterInDI(builder, envCfg)
	if err != nil {
		return
	}

	if envCfg.ENV == "test" {
		err = db.RegisterTestInDI(builder)
	} else {
		err = db.RegisterInDI(builder)
	}
	if err != nil {
		return
	}

	err = user.RegisterInDI(builder)
	if err != nil {
		return
	}

	err = langka.RegisterInDI(builder)
	if err != nil {
		return
	}

	err = webapp.RegisterInDI(builder)
	if err != nil {
		return
	}

	di = builder.Build()

	return
}
