package mailer

import (
	"context"
	"io"
)

// WithData, which wraps RawMailer in order to provide reasonable mailer.
type WithData struct {
	RawMailer     RawMailer
	HeaderFactory func(ctx context.Context, data interface{}) (header EmailHeader, err error)
	BodyFactroy   func(ctx context.Context, data interface{}, header EmailHeader) (body io.ReadCloser, err error)
}

func (mailer *WithData) SendEmail(ctx context.Context, data interface{}) (err error) {
	header, err := mailer.HeaderFactory(ctx, data)
	if err != nil {
		return
	}

	body, err := mailer.BodyFactroy(ctx, data, header)
	if err != nil {
		return
	}

	defer body.Close()

	err = mailer.RawMailer.SendEmail(ctx, header, body)
	if err != nil {
		return
	}

	return
}
