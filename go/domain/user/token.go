package user

import (
	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/util/token"
)

type RefreshToken struct {
	token.JWTStandardClaims
	Nonce  string `json:"nce"`
	UserID db.ID  `json:"uid"`
}
