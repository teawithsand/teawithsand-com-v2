package validator

import "context"

type Noop struct{}

func (*Noop) Validate(ctx context.Context, val interface{}) (err error) {
	return
}
