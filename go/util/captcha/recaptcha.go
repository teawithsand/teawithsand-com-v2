package captcha

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type RecaptchaConfig struct {
	RecaptchaPublicKey string `required:"true" split_words:"true"`
	RecaptchaSecretKey string `required:"true" split_words:"true"`
}

type RecaptchaValidator struct {
	RecaptchaConfig
	Client *http.Client
}

type recaptchaResponse struct {
	Success bool `json:"success"`
}

func (cv *RecaptchaValidator) ValidateCaptchaResponse(ctx context.Context, userResponse string) (err error) {
	data := url.Values{}
	data.Set("secret", cv.RecaptchaSecretKey)
	data.Set("response", userResponse)

	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://www.google.com/recaptcha/api/siteverify", bytes.NewBuffer([]byte(data.Encode())))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	res, err := cv.Client.Do(req)
	if err != nil {
		return
	}

	var r io.Reader = res.Body

	r = io.LimitReader(r, 1024*1024*16)

	var parsedRes recaptchaResponse
	err = json.NewDecoder(r).Decode(&parsedRes)
	if err != nil {
		return
	}

	if !parsedRes.Success {
		err = ErrInvalid
	}

	return
}
