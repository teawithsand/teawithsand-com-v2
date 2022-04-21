package webapp

import (
	"crypto/tls"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
	"github.com/teawithsand/webpage/domain/common/dikey"
	"github.com/teawithsand/webpage/domain/langka"
	"github.com/teawithsand/webpage/domain/user"
	"github.com/teawithsand/webpage/util"
	"github.com/teawithsand/webpage/util/httpext"
)

func RegisterInDI(builder util.Builder) (err error) {
	var config Config

	err = util.ReadConfig(&config)
	if err != nil {
		return
	}

	builder.Add(util.Def{
		Name: ConfigDI,
		Build: func(ctn util.DI) (interface{}, error) {
			return &config, nil
		},
	})

	builder.Add(util.Def{
		Name: RouterDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			r := mux.NewRouter()
			// TODO(teawithsand): add host before going prod, preferrably in .env
			const basePath = "/workspace/go/__dist"

			d := http.Dir(basePath)
			fs := httpext.PrecompressedHandler(d, nil)
			// TODO(teawithsand): implement serving pre-gzipped file

			homeHandler := MakeDebugHomeHandler(basePath)

			fmw := httpext.CacheMW{
				ForceDisable: true,
			}

			// TODO(teawithsand): serve gzipped content
			r.PathPrefix("/dist").Handler(
				fmw.Apply(http.StripPrefix("/dist", fs)),
			)
			r.Path("/").Methods("GET", "HEAD").Handler(homeHandler)

			var loader httpext.Loader
			var responder httpext.Responder

			err = ctn.Fill(dikey.LoaderDI, &loader)
			if err != nil {
				return
			}

			err = ctn.Fill(dikey.ResponderDI, &responder)
			if err != nil {
				return
			}

			var authMW httpext.Middleware
			err = ctn.Fill(dikey.UserAuthMiddlewareDI, &authMW)
			if err != nil {
				return
			}

			// TODO(teawithsand): make this operate on custom router, which gets pinned to master one
			// in order to log method/host not allowed requests
			brr := &baseRouteRegister{
				prefix: "/api/v1",
				r:      r,
				fac: &httpext.SimpleHandlerFactory{
					Loader:    loader,
					Responder: responder,
					ArgumentsFactory: func(r *http.Request) []reflect.Value {
						return []reflect.Value{
							reflect.ValueOf(r.Context()),
						}
					},
				},

				MW: httpext.Middlewares{
					authMW,
					&httpext.CacheMW{
						ForceDisable: true,
					},
					// httpext.LoggingMiddleware(log.Default()),
				},
			}

			// User stuff here

			var userService *user.UserService

			err = ctn.Fill(dikey.UserServiceDI, &userService)
			if err != nil {
				return
			}

			brr.registerUserRoutes(regOptions{}, userService)

			// Langka stuff here

			var langkaWordSetService *langka.WordSetService

			err = ctn.Fill(dikey.LangkaWordSetServiceDI, &langkaWordSetService)
			if err != nil {
				return
			}
			brr.registerLangkaWordSetRoutes(regOptions{}, langkaWordSetService)

			res = httpext.LoggingMiddleware(log.Default()).Apply(r)
			return
		},
	})

	builder.Add(util.Def{
		Name: HTTPServerDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			var h http.Handler
			err = ctn.Fill(RouterDI, &h)
			if err != nil {
				return
			}

			s := &http.Server{
				MaxHeaderBytes: 1024 * 8,
				Handler:        h,

				WriteTimeout: time.Second * 30,
				ReadTimeout:  time.Second * 30,
				IdleTimeout:  time.Second * 15,
			}

			res = s
			return
		},
	})

	builder.Add(util.Def{
		Name: SecureHTTPServerDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			cfg := &tls.Config{
				MinVersion: tls.VersionTLS13,
			}

			var h http.Handler
			err = ctn.Fill(RouterDI, &h)
			if err != nil {
				return
			}

			s := &http.Server{
				MaxHeaderBytes: 1024 * 8,
				Handler:        h,
				TLSConfig:      cfg,

				WriteTimeout: time.Second * 30,
				ReadTimeout:  time.Second * 30,
				IdleTimeout:  time.Second * 15,
			}

			res = s
			return
		},
	})

	builder.Add(util.Def{
		Name: RunnerDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			res = util.RunnerFunc(func() (err error) {
				return runHttp(ctn)
			})
			return
		},
	})

	builder.Add(util.Def{
		Name: DevRunnerDI,
		Build: func(ctn di.Container) (res interface{}, err error) {
			res = util.RunnerFunc(func() (err error) {
				return runHttp(ctn)
			})
			return
		},
	})

	return
}
