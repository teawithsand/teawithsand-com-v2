package explain

import "net/http"

// explainers often do not use di dependencies
// and itself are very simple
// so it's simpler to make them globals

var globalExplainers []Explainer

var globalFinalExplainer Explainer = &FinalExplainer{
	EE: &ExplainedError{
		Status:        http.StatusInternalServerError,
		DebugMessage:  "Internal Server Error",
		MessageKey:    "common.error.http.internal_server_error",
		MessageParams: nil,
	},
}

var debugFinalExplainer Explainer = ExplainerFunc(func(inErr error) (ee *ExplainedError, err error) {
	ee = &ExplainedError{
		Status:        http.StatusInternalServerError,
		DebugMessage:  inErr.Error(),
		MessageKey:    "common.error.http.internal_server_error",
		MessageParams: nil,
	}
	return
})

// TODO(teawithsand): remove global stuff, since now explainers are in separate package, which removes sense of existence of global explainers

func AddGlobal(e Explainer) {
	globalExplainers = append(globalExplainers, e)
}

func SetDebug() {
	globalFinalExplainer = debugFinalExplainer
}

func GetGlobal() Explainer {
	exps := make(Explainers, 0, len(globalExplainers)+1)
	exps = append(exps, globalExplainers...)
	exps = append(exps, globalFinalExplainer)
	return exps
}
