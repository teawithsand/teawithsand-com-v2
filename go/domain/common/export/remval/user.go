package remval

import (
	"github.com/teawithsand/ndlvr/builder"
	"github.com/teawithsand/webpage/util/regexutil"
)

var usernameValidatorRules = builder.NewFieldBuilder().
	AddRequired().
	AddMinLength(4).
	AddMaxLength(32).
	MustAddLikeRule(regexutil.AlphabethRegexp(regexutil.DefaultNameAlphabeth))

var passwordValidationRules = builder.NewFieldBuilder().
	AddRequired().
	AddMinLength(8).
	AddMaxLength(64)

var emailValidationRules = builder.NewFieldBuilder().
	AddRequired().
	AddMinLength(3).
	AddMaxLength(256).
	AddSimpleRule("email")

var UserRegistrationValidationRules = builder.NewBuilder().
	AddFieldBuilder("username", usernameValidatorRules).
	AddFieldBuilder("email", emailValidationRules).
	MustBuild()

var UserConfirmRegistrationRules = builder.NewBuilder().
	AddFieldBuilder("password", passwordValidationRules).
	MustBuild()

var UserChangePasswordValidationRules = builder.NewBuilder().
	AddFieldBuilder("password", passwordValidationRules).
	MustBuild()

var UserChangeEmailValidationRules = builder.NewBuilder().
	AddFieldBuilder("email", emailValidationRules).
	MustBuild()
