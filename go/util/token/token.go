package token

import "context"

type Manager interface {
	ParseToken(ctx context.Context, data string, res interface{}) (err error)
	IssueToken(ctx context.Context, data interface{}) (token string, err error)
}
