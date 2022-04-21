package captcha

import "context"

type TestValidator struct{}

func (cv *TestValidator) ValidateCaptchaResponse(ctx context.Context, res string) (err error) {
	if res == "iamnotarobot" {
		return
	}
	err = ErrInvalid
	return
}
