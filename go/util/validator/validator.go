package validator

import (
	"context"
)

type Validator interface {
	Validate(ctx context.Context, data interface{}) (err error)
}

type ValidatorFunc func(ctx context.Context, data interface{}) (err error)

func (f ValidatorFunc) Validate(ctx context.Context, data interface{}) (err error) {
	return f(ctx, data)
}
