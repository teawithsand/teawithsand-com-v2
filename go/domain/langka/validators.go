package langka

import (
	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/webpage/domain/common/export/remval"
	"github.com/teawithsand/webpage/util/validator"
)

type Validators struct {
	WordSetCreateDataValidator ndlvr.Validator
}

var DefaultValidators = Validators{
	WordSetCreateDataValidator: validator.MustNDLVRCompile(remval.LangkaWordSetCreateDataValidationRules),
}
