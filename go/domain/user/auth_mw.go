package user

import (
	"context"
	"net/http"
	"regexp"

	"github.com/teawithsand/webpage/util"
	"github.com/teawithsand/webpage/util/cext"
	"github.com/teawithsand/webpage/util/dbutil"
	"github.com/teawithsand/webpage/util/token"
	"github.com/teawithsand/webpage/util/voter"
)

type AuthMW struct {
	TokenManager   token.Manager
	UserRepository UserRepository
	Checker        voter.Checker
}

var prefixRegex = regexp.MustCompile("(?i)^bearer ")

func (amw *AuthMW) ParseHeader(ctx context.Context, header string) (user *User, err error) {
	if !prefixRegex.MatchString(header) {
		err = ErrInvalidAuthHeader
		return
	}
	token := header[len("bearer "):]

	return amw.ParseToken(ctx, token)
}

func (amw *AuthMW) ParseToken(ctx context.Context, token string) (user *User, err error) {
	var parsedToken RefreshToken
	err = amw.TokenManager.ParseToken(ctx, token, &parsedToken)
	if err != nil {
		return
	}

	var innerUser User
	ok, err := amw.UserRepository.FindOne(ctx, dbutil.NewIDQuery(parsedToken.UserID), &innerUser)
	if err != nil {
		return
	}

	if !ok {
		err = ErrTokenUserNotFound
		return
	}

	if !util.SafeStringEquals(innerUser.TokenAuthData.Nonce, parsedToken.Nonce) {
		err = ErrTokenUserNonceMismatch
		return
	}

	err = amw.Checker.Check(ctx, voter.Voting{
		Operation: VoteUseApp,
		Object:    user,
	})
	if err != nil {
		return
	}

	user = &innerUser

	return
}

func (amw *AuthMW) Apply(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		var user *User
		header := r.Header.Get(AuthHeaderName)
		if len(header) > 0 {
			user, err = amw.ParseHeader(r.Context(), header)
		}

		newContext := r.Context()
		if err != nil {
			newContext = cext.PutError(newContext, err)
		} else {
			newContext = PutUser(newContext, user)
		}

		r = r.WithContext(newContext)
		h.ServeHTTP(w, r)
	})
}
