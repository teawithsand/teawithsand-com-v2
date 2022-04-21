package common

import (
	"net/http"
	"time"

	"github.com/sarulabs/di"
	"github.com/teawithsand/webpage/domain/common/dikey"
	"github.com/teawithsand/webpage/domain/common/export/remexpl"
	"github.com/teawithsand/webpage/domain/common/permcheck"
	"github.com/teawithsand/webpage/util"
	"github.com/teawithsand/webpage/util/captcha"
	"github.com/teawithsand/webpage/util/env"
	"github.com/teawithsand/webpage/util/explain"
	"github.com/teawithsand/webpage/util/httpext"
	"github.com/teawithsand/webpage/util/mailer"
	"github.com/teawithsand/webpage/util/phash"
	"github.com/teawithsand/webpage/util/slug"
)

func RegisterInDI(builder util.Builder, cfg env.Config) (err error) {
	httpClient := http.Client{
		Timeout: time.Second * 15,
	}

	builder.Add(di.Def{
		Name: dikey.EnvDI,
		Build: func(ctn di.Container) (interface{}, error) {
			return &cfg, nil
		},
	})

	builder.Add(di.Def{
		Name: dikey.SluggifierDI,
		Build: func(ctn di.Container) (interface{}, error) {
			return &slug.DefaultSluggifier{}, nil
		},
	})

	builder.Add(di.Def{
		Name: dikey.CheckerDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			res, err = permcheck.MakeChecker()
			if err != nil {
				return
			}
			return
		},
	})

	builder.Add(di.Def{
		Name: dikey.CaptchaValidatorDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			var env *env.Config
			err = ctn.Fill(dikey.EnvDI, &env)
			if err != nil {
				return
			}

			if env.ENV == "test" {
				res = &captcha.TestValidator{}
			} else {
				var captchaCfg captcha.RecaptchaConfig
				err = util.ReadConfig(&captchaCfg)
				if err != nil {
					return
				}

				res = &captcha.RecaptchaValidator{
					RecaptchaConfig: captchaCfg,
					Client:          &httpClient,
				}
			}
			return
		},
	})

	builder.Add(di.Def{
		Name: dikey.PasswordHasherDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			// TODO(teawithsand): change that before going prod

			var env *env.Config
			err = ctn.Fill(dikey.EnvDI, &env)
			if err != nil {
				return
			}

			// TODO(teawithsand): run it only when env is test
			res = &phash.NoopHasher{}
			return
		},
	})

	builder.Add(di.Def{
		Name: dikey.RawMailerDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			// TODO(teawithsand): change that before going prod
			res = &mailer.LoggingRaw{}
			return
		},
	})

	builder.Add(di.Def{
		Name: dikey.ResponderDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			var env *env.Config
			err = ctn.Fill(dikey.EnvDI, &env)
			if err != nil {
				return
			}

			if env.ENV == "dev" || env.ENV == "test" {
				explain.SetDebug()
			}

			explainer := explain.Responder{
				InnerResponder: &httpext.JSONResponder{},
			}

			err = ctn.Fill(dikey.ExplainerDI, &explainer.Explainer)
			if err != nil {
				return
			}

			res = &explainer
			return
		},
	})

	builder.Add(di.Def{
		Name: dikey.LoaderDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			res = &httpext.JSONLoader{}
			return
		},
	})

	builder.Add(di.Def{
		Name: dikey.ExplainerDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			remexpl.InitExplainers()

			res = explain.GetGlobal()
			return
		},
	})

	return
}
