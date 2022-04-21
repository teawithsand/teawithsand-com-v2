package domtestutil

import (
	"context"

	"github.com/teawithsand/webpage/domain"
	"github.com/teawithsand/webpage/domain/common/dikey"
	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/domain/langka"
	"github.com/teawithsand/webpage/domain/user"
	"github.com/teawithsand/webpage/util"
	"github.com/teawithsand/webpage/util/testutil"
	"github.com/teawithsand/webpage/util/token"
)

type TestDI struct {
	DI util.DI

	Ctx context.Context

	LangkaWordSetService    *langka.WordSetService
	LangkaWordSetRepository langka.WordSetRepository

	UserService         *user.UserService
	UserRepository      user.UserRepository
	AuthMW              *user.AuthMW
	RefreshTokenManager token.Manager

	DB *db.TypedDB
}

func (tdi *TestDI) Close() (err error) {
	return tdi.DI.Delete()
}
func (tdi *TestDI) MustClose() {
	err := tdi.Close()
	if err != nil {
		panic(err)
	}
}

func MustConstructTestDI() (tdi *TestDI) {
	tdi, err := ConstructTestDI()
	if err != nil {
		panic(err)
	}
	return
}

func ConstructTestDI() (tdi *TestDI, err error) {
	testutil.SetTestENV()

	di, err := domain.ConstructDI()
	if err != nil {
		return
	}

	tdi = &TestDI{
		DI:  di,
		Ctx: context.Background(),
	}

	err = di.Fill(dikey.UserServiceDI, &tdi.UserService)
	if err != nil {
		return
	}

	err = di.Fill(dikey.UserRepositoryDI, &tdi.UserRepository)
	if err != nil {
		return
	}

	err = di.Fill(dikey.UserAuthMiddlewareDI, &tdi.AuthMW)
	if err != nil {
		return
	}

	err = di.Fill(dikey.UserRefreshTokenManagerDI, &tdi.RefreshTokenManager)
	if err != nil {
		return
	}

	err = di.Fill(dikey.DBDatabaseDI, &tdi.DB)
	if err != nil {
		return
	}

	err = di.Fill(dikey.LangkaWordSetServiceDI, &tdi.LangkaWordSetService)
	if err != nil {
		return
	}

	err = di.Fill(dikey.LangkaWordSetRepositoryDI, &tdi.LangkaWordSetRepository)
	if err != nil {
		return
	}

	return
}
