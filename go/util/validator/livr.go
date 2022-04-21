package validator

import (
	"context"

	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/ndlvr/builtin"
	"github.com/teawithsand/ndlvr/value"
)

// Note: these are deprecated, once we will move to NDLVR

var bf = builtin.MakeBuiltinFactory()

func MustNDLVRCompile(rules interface{}) ndlvr.Validator {
	opts := ndlvr.Options{
		ValidationFactory: bf,
	}

	source, ok := rules.(ndlvr.TopRulesSource)
	if !ok {
		source = ndlvr.RulesMap(rules.(map[string]interface{}))
	}

	ng, err := opts.NewEngine(context.Background(), source)
	if err != nil {
		panic(err)
	}

	return &ndlvr.DefaultValidator{
		Wrapper: &value.DefaultWrapper{
			UseJSONNames: true,
		},
		Engine: ng,
	}
}
