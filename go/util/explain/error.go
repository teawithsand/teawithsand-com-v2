package explain

import (
	"net/http"

	"github.com/teawithsand/webpage/util/httpext"
)

// Error, which is explained for end user.
// All errors should be translated to this form before passing into responder in order to send them to the client.
type ExplainedError struct {
	Status        int               `json:"status"`
	DebugMessage  string            `json:"debugMessage"`
	MessageKey    string            `json:"messageKey"`
	MessageParams map[string]string `json:"messageParams"`
}

func (ee *ExplainedError) GetStatus() int {
	if ee.Status == 0 {
		return http.StatusInternalServerError
	}
	return ee.Status
}

var _ httpext.StatusResponse = &ExplainedError{}
