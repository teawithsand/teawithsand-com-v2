package explain

import "errors"

// TODO(teawithsand): use nil EE instaead
var ErrCantExplain = errors.New("util/explain: this explainer can't explain this error")

// Explainer is something, which translates error into explained one.
type Explainer interface {
	ExplainError(inErr error) (ee *ExplainedError, err error)
}

type ExplainerFunc func(inErr error) (ee *ExplainedError, err error)

func (f ExplainerFunc) ExplainError(inErr error) (ee *ExplainedError, err error) {
	return f(inErr)
}

// Groups multiple explainers in order to return explainer, which tries to explain error with all explainers it consists of.
type Explainers []Explainer

func (exps Explainers) ExplainError(inErr error) (ee *ExplainedError, err error) {
	for _, e := range exps {
		ee, err = e.ExplainError(inErr)
		if errors.Is(err, ErrCantExplain) {
			continue
		} else if err != nil {
			return
		} else {
			return
		}
	}
	ee = &ExplainedError{}
	err = ErrCantExplain
	return
}

// FinalExplainer is explainer, which just always returns error it's given.
type FinalExplainer struct {
	EE *ExplainedError
}

func (fe *FinalExplainer) ExplainError(inErr error) (ee *ExplainedError, err error) {
	ee = fe.EE
	return
}
