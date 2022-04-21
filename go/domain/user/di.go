package user

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/sarulabs/di"
	"github.com/teawithsand/webpage/domain/common/dikey"
	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/util"
	"github.com/teawithsand/webpage/util/mailer"
	"github.com/teawithsand/webpage/util/token"
)

func RegisterInDI(builder util.Builder) (err error) {
	var config Config
	err = util.ReadConfig(&config)
	if err != nil {
		return
	}

	builder.Add(util.Def{
		Name: dikey.UserConfigDI,
		Build: func(ctn di.Container) (interface{}, error) {
			return &config, err
		},
	})

	builder.Add(util.Def{
		Name: dikey.UserRepositoryDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			var tdb *db.TypedDB
			err = ctn.Fill(dikey.DBDatabaseDI, &tdb)
			if err != nil {
				return
			}

			res = NewUserRepository(tdb.Users)
			return
		},
	})

	builder.Add(util.Def{
		Name: dikey.UserRegistrationsRepositoryDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			var tdb *db.TypedDB
			err = ctn.Fill(dikey.DBDatabaseDI, &tdb)
			if err != nil {
				return
			}

			res = NewRegistrationRepository(tdb.Registrations)
			return
		},
	})

	builder.Add(util.Def{
		Name: dikey.UserServiceDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			var service UserService

			err = ctn.Fill(dikey.UserRepositoryDI, &service.UserRepository)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.UserRegistrationsRepositoryDI, &service.RegistrationRepository)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.CaptchaValidatorDI, &service.CaptchaValidator)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.UserRegisterUserMailerDI, &service.RegistrationMailer)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.CheckerDI, &service.Checker)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.PasswordHasherDI, &service.PasswordHasher)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.UserRegistrationTokenGeneratorDI, &service.RegistrationTokenGenerator)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.UserRefreshTokenNonceGeneratorDI, &service.RefreshTokenNonceGenerator)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.UserPasswordResetTokenGeneratorDI, &service.PasswordResetTokenGenerator)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.UserRefreshTokenManagerDI, &service.RefreshTokenManager)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.UserEmailConfirmTokenGenerator, &service.EmailConfirmTokenGenerator)
			if err != nil {
				return
			}

			// for now no di for this value
			service.Validators = DefaultValidators

			res = &service
			return
		},
	})

	builder.Add(util.Def{
		Name: dikey.UserRegisterUserMailerDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			var mlr mailer.RawMailer
			err = ctn.Fill(dikey.RawMailerDI, &mlr)
			if err != nil {
				return
			}

			res = &mailer.WithData{
				RawMailer: mlr,
				HeaderFactory: func(ctx context.Context, data interface{}) (header mailer.EmailHeader, err error) {
					tdata := data.(RegisterMailData)
					header = mailer.EmailHeader{
						Subject: "teawithsand.com - registration",
						From:    "noreply@teawithsand.com",
						To:      []string{tdata.Email},
					}
					return
				},
				BodyFactroy: func(ctx context.Context, data interface{}, header mailer.EmailHeader) (body io.ReadCloser, err error) {
					tdata := data.(RegisterMailData)
					text := fmt.Sprintf("data: %#+v\nheader: %+#v\n", tdata, header)
					body = io.NopCloser(bytes.NewBufferString(text))
					return
				},
			}
			return
		},
	})

	builder.Add(di.Def{
		Name: dikey.UserRegistrationTokenGeneratorDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			res = &util.HexNonceGenerator{
				BytesLength: 32,
			}
			return
		},
	})

	builder.Add(di.Def{
		Name: dikey.UserEmailConfirmTokenGenerator,
		Build: func(ctn di.Container) (res interface{}, err error) {
			res = &util.HexNonceGenerator{
				BytesLength: 32,
			}
			return
		},
	})

	builder.Add(di.Def{
		Name: dikey.UserRefreshTokenNonceGeneratorDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			res = &util.HexNonceGenerator{
				BytesLength: 16,
			}
			return
		},
	})

	builder.Add(di.Def{
		Name: dikey.UserPasswordResetTokenGeneratorDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			res = &util.HexNonceGenerator{
				BytesLength: 32,
			}
			return
		},
	})

	builder.Add(di.Def{
		Name: dikey.UserRefreshTokenManagerDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			res = &token.JWTManager{
				SecretKey: config.JWTRefreshTokenSecretKey,
			}
			return
		},
	})

	builder.Add(di.Def{
		Name: dikey.UserAuthMiddlewareDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			var mw AuthMW

			err = ctn.Fill(dikey.UserRefreshTokenManagerDI, &mw.TokenManager)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.CheckerDI, &mw.Checker)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.UserRepositoryDI, &mw.UserRepository)
			if err != nil {
				return
			}

			res = &mw
			return
		},
	})

	return
}
