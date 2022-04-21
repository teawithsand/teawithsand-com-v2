package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/domain/domtestutil"
	"github.com/teawithsand/webpage/domain/user"
	"github.com/teawithsand/webpage/util/testutil"
)

func TestAuthMW(t *testing.T) {
	testutil.SetTestENV()

	t.Run("malformed_ok", func(t *testing.T) {
		tdi := domtestutil.MustConstructTestDI()
		defer tdi.Close()

		_, err := tdi.AuthMW.ParseHeader(context.Background(), "Bearer asdfasdfasdf")
		if err == nil {
			t.Error("expected error")
			return
		}
	})

	t.Run("valid_ok", func(t *testing.T) {
		mockUser := user.User{
			ID: db.NewID(),
			TokenAuthData: user.TokenAuthData{
				Nonce: "asdfasdfasdfasdf",
			},
		}

		tdi := domtestutil.MustConstructTestDI()
		defer tdi.Close()

		ctx := tdi.Ctx
		ctx, closer, err := tdi.DB.MakeSession(ctx)
		if err != nil {
			t.Error(err)
			return
		}
		defer closer.Close()

		err = tdi.UserRepository.Create(ctx, mockUser)
		if err != nil {
			t.Error(err)
			return
		}

		token, err := tdi.RefreshTokenManager.IssueToken(ctx, user.RefreshToken{
			JWTStandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Second * 10).Unix(),
			},
			Nonce:  mockUser.TokenAuthData.Nonce,
			UserID: mockUser.ID,
		})
		if err != nil {
			t.Error(err)
			return
		}

		parsedUser, err := tdi.AuthMW.ParseHeader(ctx, "Bearer "+token)
		if err != nil {
			t.Error(err)
			return
		}

		if parsedUser == nil {
			t.Error("must not be nil")
		} else if parsedUser.ID != mockUser.ID {
			t.Error("id mismatch")
		}
	})
}
