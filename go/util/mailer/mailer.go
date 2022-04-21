package mailer

import (
	"context"
	"io"
)

type EmailHeader struct {
	From    string
	To      []string
	Subject string
}

type RawMailer interface {
	SendEmail(ctx context.Context, header EmailHeader, content io.Reader) (err error)
}

type Mailer interface {
	SendEmail(ctx context.Context, data interface{}) (err error)
}
