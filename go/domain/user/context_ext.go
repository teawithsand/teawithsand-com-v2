package user

import (
	"context"

	"github.com/teawithsand/webpage/util/cext"
)

var userVar = cext.ContextVar{
	Name: "user_user",
}

func PutUser(ctx context.Context, user *User) context.Context {
	return userVar.Put(ctx, user)
}

func GetUser(ctx context.Context) (user *User) {
	rawRes := userVar.Get(ctx)
	if rawRes != nil {
		user = rawRes.(*User)
	}
	return
}
