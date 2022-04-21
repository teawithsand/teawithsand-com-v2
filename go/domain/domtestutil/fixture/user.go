package fixture

import (
	"time"

	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/domain/user"
)

var MockUserOne = user.User{
	ID:         db.NewID(),
	PublicName: "userone",
	TokenAuthData: user.TokenAuthData{
		Nonce: "somenoncegoeshere",
	},
	Email: user.EmailData{
		Email:             "userone@example.com",
		EmailConfirmedAt:  time.Now(),
		EmailConfirmToken: "",
	},
	Lifecycle: user.UserLifecycle{
		CreatedAt: time.Now(),
	},
	NativeAuthData: user.NativeAuthData{
		PasswordHash: "useronepassword",
	},
}

var MockUserTwo = user.User{
	ID:         db.NewID(),
	PublicName: "usertwo",
	TokenAuthData: user.TokenAuthData{
		Nonce: "someothernoncegoeshere",
	},
	Email: user.EmailData{
		Email:             "usertwo@example.com",
		EmailConfirmedAt:  time.Now(),
		EmailConfirmToken: "",
	},
	Lifecycle: user.UserLifecycle{
		CreatedAt: time.Now(),
	},
	NativeAuthData: user.NativeAuthData{
		PasswordHash: "usertwopassword",
	},
}
