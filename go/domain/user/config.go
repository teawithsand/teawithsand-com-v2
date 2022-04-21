package user

import "net/textproto"

var AuthHeaderName = textproto.CanonicalMIMEHeaderKey("WWW-Authenticate")

type Config struct {
	JWTRefreshTokenSecretKey string `required:"true" split_words:"true"`
	JWTAuthTokenSecretKey    string `required:"true" split_words:"true"`
}
