package explain

import "errors"

func MakeIsExplainer(err error, ee ExplainedError) *IsExplainer {
	ee.DebugMessage = err.Error()
	return &IsExplainer{
		Error: err,
		EE:    &ee,
	}
}

// Explainer, which compares error given to pattern using errors.Is function
type IsExplainer struct {
	EE    *ExplainedError
	Error error
}

func (e *IsExplainer) ExplainError(inErr error) (ee *ExplainedError, err error) {
	if !errors.Is(inErr, e.Error) {
		err = ErrCantExplain
		return
	}

	ee = e.EE
	return
}
