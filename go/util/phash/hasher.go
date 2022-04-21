package phash

import (
	"context"
	"errors"
)

var ErrPasswordMismatch = errors.New("util/hasher: password does not match its hash")

type Hasher interface {
	HashPassword(ctx context.Context, pass string) (res string, err error)
	CheckPassword(ctx context.Context, pass, hash string) (err error)
}

// TODO(teawithsand): integrate this type with crypka
