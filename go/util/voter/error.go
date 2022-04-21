package voter

import "errors"

var ErrAccessDenied = errors.New("util/voter: Voter access denied for given voting")
