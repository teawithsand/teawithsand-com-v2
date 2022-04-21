package token

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt"
)

var ErrNotJWTClaims = errors.New("util/token: data/receiver given must be jwt claims")
var ErrTokenInvalid = errors.New("util/token: provided token is not valid")

type JWTStandardClaims = jwt.StandardClaims

type JWTManager struct {
	SecretKey string
	// For now algo is fixed SHA256
}

func (tm *JWTManager) ParseToken(ctx context.Context, data string, res interface{}) (err error) {
	resClaims, ok := res.(jwt.Claims)
	if !ok {
		err = ErrNotJWTClaims
		return
	}

	token, err := jwt.ParseWithClaims(data, resClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte(tm.SecretKey), nil
	})
	if err != nil {
		err = ErrTokenInvalid
		return
	}

	if !token.Valid {
		err = ErrTokenInvalid
		return
	}

	return
}
func (tm *JWTManager) IssueToken(ctx context.Context, data interface{}) (token string, err error) {
	claimsData, ok := data.(jwt.Claims)
	if !ok {
		err = ErrNotJWTClaims
		return
	}

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsData)
	token, err = tk.SignedString([]byte(tm.SecretKey))
	return
}
