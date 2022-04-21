package captcha

import (
	"context"
	"errors"
)

var ErrInvalid = errors.New("util/captcha: invalid captcha response")

type Validator interface {
	ValidateCaptchaResponse(ctx context.Context, res string) (err error)
}
