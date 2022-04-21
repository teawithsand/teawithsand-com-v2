package explain

import (
	"context"
	"io"

	"github.com/teawithsand/webpage/util/httpext"
)

// Responder, which checks if type passed is error, and explains it if it is.
type Responder struct {
	Explainer      Explainer
	InnerResponder httpext.Responder
}

func (rs *Responder) RespondWithData(ctx context.Context, w io.Writer, res interface{}) (err error) {
	errRes, ok := res.(error)
	if ok {
		res, err = rs.Explainer.ExplainError(errRes)
		// should be unreachable code with final explainer
		if err != nil {
			panic(err)
		}
	}

	return rs.InnerResponder.RespondWithData(ctx, w, res)
}

var _ httpext.Responder = &Responder{}
