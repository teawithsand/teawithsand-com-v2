package webapp

import (
	"github.com/teawithsand/webpage/domain/user"
)

func (brr *baseRouteRegister) registerUserRoutes(opts regOptions, svc *user.UserService) {
	brr.r.Path(brr.prefix + "/user/register/init").
		Methods("POST").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithBody(svc.RegisterUser),
			opts,
		))

	brr.r.Path(brr.prefix + "/user/register/confirm").
		Methods("POST").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithBody(svc.ConfirmRegistration),
			opts,
		))

	brr.r.Path(brr.prefix + "/user/login").
		Methods("POST").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithBody(svc.LoginUser),
			opts,
		))

	brr.r.Path(brr.prefix + "/user/invalidate-tokens").
		Methods("POST").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithBody(svc.InvalidateTokens),
			opts,
		))

	brr.r.Path(brr.prefix + "/user/reset-password/init").
		Methods("POST").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithBody(svc.InitPasswordReset),
			opts,
		))

	brr.r.Path(brr.prefix + "/user/reset-password/finalize").
		Methods("POST").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithBody(svc.PasswordReset),
			opts,
		))

	brr.r.Path(brr.prefix + "/user/change-password").
		Methods("POST").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithBody(svc.ChangePassword),
			opts,
		))

	brr.r.Path(brr.prefix + "/user/change-email/init").
		Methods("POST").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithBody(svc.ChangeEmail),
			opts,
		))

	brr.r.Path(brr.prefix + "/user/change-email/finalize").
		Methods("POST").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithBody(svc.ConfirmEmail),
			opts,
		))

	brr.r.Path(brr.prefix+"/user/profile/secret").
		Methods("GET", "HEAD").
		Handler(brr.apply(
			brr.fac.MustMakeHandlerWithQuery(svc.GetUserSecretProfile),
			opts,
		))
}
