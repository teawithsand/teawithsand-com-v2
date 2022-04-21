package mailer

import (
	"context"
	"io"
	"log"
)

type LoggingRaw struct {
	*log.Logger
}

func (lm *LoggingRaw) SendEmail(ctx context.Context, header EmailHeader, content io.Reader) (err error) {
	data, err := io.ReadAll(content)
	if err != nil {
		return
	}

	l := lm.Logger
	if l == nil {
		l = log.Default()
	}

	l.Println("Mock email sent header:", header, "Body: ", string(data))
	return
}

type Logging struct {
	*log.Logger
}

func (lm *Logging) SendEmail(ctx context.Context, data interface{}) (err error) {
	l := lm.Logger
	if l == nil {
		l = log.Default()
	}

	l.Println("Mock email sent", data)
	return
}
