package langka

import "errors"

var ErrWordSetNotFound = errors.New("domain/langka: given word set does not exist")
var ErrWordSetNameInUse = errors.New("domain/langka: given word set name is already in use")
