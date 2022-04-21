package webapp

import (
	"github.com/teawithsand/webpage/domain/langka"
)

func (brr *baseRouteRegister) registerLangkaWordSetRoutes(opts regOptions, svc *langka.WordSetService) {
	brr.r.Path(brr.prefix + "/langka/word-set").
		Methods("POST").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithBody(svc.CreateWordSet),
			opts,
		))

	brr.r.Path(brr.prefix+"/langka/word-set").
		Methods("PATCH", "PUT").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithBody(svc.EditWordSet),
			opts,
		))

	brr.r.Path(brr.prefix + "/langka/word-set").
		Methods("DELETE").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithBody(svc.DeleteWordSet),
			opts,
		))

	brr.r.Path(brr.prefix + "/langka/word-set/publish").
		Methods("POST").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithBody(svc.PublishWordSet),
			opts,
		))

	brr.r.Path(brr.prefix+"/langka/word-set/public").
		Methods("GET", "HEAD").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithQuery(svc.GetPublicWordSet),
			opts,
		))

	brr.r.Path(brr.prefix+"/langka/word-set/secret").
		Methods("GET", "HEAD").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithQuery(svc.GetSecretWordSet),
			opts,
		))

	brr.r.Path(brr.prefix+"/langka/word-set/list/public").
		Methods("GET", "HEAD").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithQuery(svc.GetPublicWordSetList),
			opts,
		))

	brr.r.Path(brr.prefix+"/langka/word-set/list/owned").
		Methods("GET", "HEAD").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithQuery(svc.GetOwnerdWordSetList),
			opts,
		))
}
