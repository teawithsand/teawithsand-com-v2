package phash

import "context"

type NoopHasher struct{}

func (nph *NoopHasher) HashPassword(ctx context.Context, pass string) (res string, err error) {
	res = pass
	return
}

func (nph *NoopHasher) CheckPassword(ctx context.Context, pass, hash string) (err error) {
	if pass != hash {
		err = ErrPasswordMismatch
		return
	}
	return
}
