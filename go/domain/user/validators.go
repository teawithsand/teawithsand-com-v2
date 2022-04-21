package user

import (
	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/webpage/domain/common/export/remval"
	"github.com/teawithsand/webpage/util/validator"
)

type Validators struct {
	RegistrationCreateDataValidator ndlvr.Validator
	ChangePasswordDataValidator     ndlvr.Validator
	ChangeEmailDataValidator        ndlvr.Validator
	ConfirmRegistrationValidator    ndlvr.Validator
}

var DefaultValidators = Validators{
	RegistrationCreateDataValidator: validator.MustNDLVRCompile(remval.UserRegistrationValidationRules),
	ChangePasswordDataValidator:     validator.MustNDLVRCompile(remval.UserChangePasswordValidationRules),
	ChangeEmailDataValidator:        validator.MustNDLVRCompile(remval.UserChangeEmailValidationRules),
	ConfirmRegistrationValidator:    validator.MustNDLVRCompile(remval.UserConfirmRegistrationRules),
}
