package validator

import "context"

// Validator, which returns error when at least one of validators return error.
type AndValidators []Validator

func (cv AndValidators) Validate(ctx context.Context, data interface{}) (err error) {
	for _, v := range cv {
		err = v.Validate(ctx, data)
		if err != nil {
			return
		}
	}
	return
}

// Validator, which returns nil when at least one of validators return nil.
// Otherwise returns last validator's non-nil error.
type OrValidators []Validator

func (cv OrValidators) Validate(ctx context.Context, data interface{}) (err error) {
	for _, v := range cv {
		err = v.Validate(ctx, data)
		if err == nil {
			return
		}
	}
	return
}
