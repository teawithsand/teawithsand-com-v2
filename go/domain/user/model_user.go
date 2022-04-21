package user

import (
	"time"

	"github.com/teawithsand/webpage/domain/db"
	"go.mongodb.org/mongo-driver/bson"
)

type EmailData struct {
	Email             string
	EmailConfirmedAt  time.Time
	EmailConfirmToken string
}

type NativeAuthData struct {
	PasswordHash       string
	PasswordResetToken string
}

type TokenAuthData struct {
	Nonce string
}

// User model stored in db.
type User struct {
	ID         db.ID `bson:"_id"`
	PublicName string

	TokenAuthData TokenAuthData

	Email          EmailData
	Lifecycle      UserLifecycle
	NativeAuthData NativeAuthData
}

func (u *User) GetReference() (ref UserReference) {
	ref.Load(u)
	return
}

type UserReference struct {
	ID         db.ID `bson:"_id"`
	PublicName string
}
type UserReferenceProjection struct {
	ID         db.ID  `json:"id" bson:"_id"`
	PublicName string `json:"publicName"`
}

func (proj *UserReferenceProjection) Load(ref *UserReference) {
	proj.ID = ref.ID
	proj.PublicName = ref.PublicName
}

func (ref *UserReference) Load(user *User) {
	ref.ID = user.ID
	ref.PublicName = user.PublicName
}

type UserLifecycle struct {
	CreatedAt time.Time
	LockedAt  time.Time
	DeletedAt time.Time
}

func NewPublicNameUserQuery(name string) interface{} {
	return bson.M{
		"publicname": name,
	}
}

func NewEmailUserQuery(email string) interface{} {
	return bson.M{
		"email": bson.M{
			"email": email,
		},
	}
}

func NewTokenNonceUserUpdate(nonce string) interface{} {
	return bson.M{
		"$set": bson.M{
			"tokenauthdata": bson.M{
				"nonce": nonce,
			},
		},
	}
}

func NewPasswordAndTokensAndNoncesUserUpdate(passwordHash, tokenNonce, passwordResetToken string) interface{} {
	return bson.M{
		"$set": bson.M{
			"tokenauthdata": bson.M{
				"nonce": tokenNonce,
			},
			"nativeauthdata": bson.M{
				"passwordhash":       passwordHash,
				"passwordresettoken": passwordResetToken,
			},
		},
	}
}

func NewEmailAndNonceUserUpdate(email, emailConfirmToken string) interface{} {
	return bson.M{
		"$set": bson.M{
			"email": EmailData{
				Email:             email,
				EmailConfirmedAt:  time.Time{},
				EmailConfirmToken: emailConfirmToken,
			},
		},
	}
}

func NewEmailConfirmedUserUpdate(confirmedAt time.Time) interface{} {
	return bson.M{
		"$set": bson.M{
			"email": bson.M{
				"emailconfirmedat": confirmedAt,
			},
		},
	}
}

type PublicUserProjection struct {
	ID         db.ID  `json:"id"`
	PublicName string `json:"publicName"`
}

func (proj *PublicUserProjection) Load(user *User) {
	proj.ID = user.ID
	proj.PublicName = user.PublicName
}

type SecretUserProjection struct {
	ID               db.ID      `json:"id"`
	PublicName       string     `json:"publicName"`
	Email            string     `json:"email"`
	RegisteredAt     time.Time  `json:"registeredAt"`
	EmailConfirmedAt *time.Time `json:"emailConfirmedAt"`
}

func (proj *SecretUserProjection) Load(user *User) {
	proj.ID = user.ID
	proj.PublicName = user.PublicName
	proj.Email = user.Email.Email
	proj.RegisteredAt = user.Lifecycle.CreatedAt

	// TODO(teawithsand): move this to some null time util with custom JSON marshaler
	if !user.Email.EmailConfirmedAt.IsZero() {
		proj.EmailConfirmedAt = &user.Email.EmailConfirmedAt
	}
}
