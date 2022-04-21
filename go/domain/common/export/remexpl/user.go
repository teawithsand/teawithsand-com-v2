package remexpl

import (
	"net/http"

	"github.com/teawithsand/webpage/domain/user"
	"github.com/teawithsand/webpage/util/explain"
)

func registerUserExplainers(explainers explain.Explainers) (res explain.Explainers, err error) {
	explainers = append(explainers, explain.MakeIsExplainer(
		user.ErrLoginInUse,
		explain.ExplainedError{
			Status:     http.StatusUnprocessableEntity,
			MessageKey: "user.error.login_in_use",
		},
	))
	explainers = append(explainers, explain.MakeIsExplainer(
		user.ErrEmailInUse,
		explain.ExplainedError{
			Status:     http.StatusUnprocessableEntity,
			MessageKey: "user.error.email_in_use",
		},
	))

	var tokenError = explain.ExplainedError{
		Status:     http.StatusForbidden,
		MessageKey: "user.error.auth_token_invalid",
	}

	explainers = append(explainers, explain.MakeIsExplainer(
		user.ErrInvalidAuthHeader,
		tokenError,
	))

	explainers = append(explainers, explain.MakeIsExplainer(
		user.ErrTokenUserNotFound,
		tokenError,
	))

	explainers = append(explainers, explain.MakeIsExplainer(
		user.ErrTokenUserNonceMismatch,
		tokenError,
	))

	explainers = append(explainers, explain.MakeIsExplainer(
		user.ErrUserForLoginNotFound,
		explain.ExplainedError{
			Status:     http.StatusNotFound,
			MessageKey: "user.error.not_found_for_login",
		},
	))

	explainers = append(explainers, explain.MakeIsExplainer(
		user.ErrUserAlreadyRegistered,
		explain.ExplainedError{
			Status:     http.StatusUnprocessableEntity,
			MessageKey: "user.error.user_for_registration_already_registered",
		},
	))

	res = explainers
	return
}
