package webapp

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teawithsand/webpage/util/httpext"
)

type baseRouteRegister struct {
	prefix string

	r   *mux.Router
	fac *httpext.SimpleHandlerFactory

	MW httpext.Middleware
}

func (brr *baseRouteRegister) apply(h http.Handler, opts regOptions) http.Handler {
	if brr.MW != nil {
		h = brr.MW.Apply(h)
	}
	h = opts.Apply(h)

	return h
}

type regOptions struct {
	MW httpext.Middleware
}

func (ro *regOptions) Apply(h http.Handler) http.Handler {
	if ro == nil || ro.MW == nil {
		return h
	}

	return ro.MW.Apply(h)
}
