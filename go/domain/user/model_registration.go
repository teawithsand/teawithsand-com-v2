package user

import (
	"context"
	"time"

	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/util"
	"github.com/teawithsand/webpage/util/cext"
	"go.mongodb.org/mongo-driver/bson"
)

type Registration struct {
	ID db.ID `bson:"_id"`

	PublicName string
	Email      string

	Token string

	CreatedAt time.Time
}

func (reg *Registration) MakeUser(
	ctx context.Context,
	tokenNonceGenerator util.NonceGenerator,
	passwordResetTokenGenerator util.NonceGenerator,
	phash string,
) (user *User, err error) {
	tokenNonce, err := tokenNonceGenerator.GenerateNonce(ctx)
	if err != nil {
		return
	}

	passwordResetToken, err := passwordResetTokenGenerator.GenerateNonce(ctx)
	if err != nil {
		return
	}

	user = &User{
		ID:         db.NewID(),
		PublicName: reg.PublicName,
		TokenAuthData: TokenAuthData{
			Nonce: tokenNonce,
		},
		Email: EmailData{
			Email:             reg.Email,
			EmailConfirmedAt:  cext.Now(ctx),
			EmailConfirmToken: "",
		},
		Lifecycle: UserLifecycle{
			CreatedAt: cext.Now(ctx),
		},
		NativeAuthData: NativeAuthData{
			PasswordHash:       phash,
			PasswordResetToken: passwordResetToken,
		},
	}
	return
}

func NewEmailRegistrationQuery(email string) interface{} {
	return bson.M{
		"email": email,
	}
}

func NewTokenRegistrationQuery(token string) interface{} {
	return bson.M{
		"token": token,
	}
}
