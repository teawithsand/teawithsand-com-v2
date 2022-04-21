package remval

import (
	"github.com/teawithsand/ndlvr/builder"
	"github.com/teawithsand/webpage/domain/common/export/remdefines"
	"github.com/teawithsand/webpage/util/regexutil"
)

var langkaWordSetNameValidationRules = builder.NewFieldBuilder().
	AddRequired().
	AddMinLength(3).
	AddMaxLength(512).
	MustAddLikeRule(regexutil.TrimmedRegexp)

var langkaWordValidationRules = builder.NewFieldBuilder().
	AddRequired().
	MustAddLikeRule(regexutil.AlphabethRegexp(regexutil.ExtendedNameAlphabeth)).
	MustAddLikeRule(regexutil.TrimmedRegexp).
	AddMinLength(1).
	AddMaxLength(128)

var langkaLanguageNameValidationRules = builder.NewFieldBuilder().
	AddRequired().
	MustAddLikeRule(regexutil.EnumRegexp(remdefines.LanguageCodes))

var langkaWordTupleDescriptionValidationRules = builder.NewFieldBuilder().
	MustAddLikeRule(regexutil.AlphabethRegexp(regexutil.DefaultDescriptionAlphabeth)).
	AddMaxLength(1024 * 16)

var langkaWordSetDescriptionValidationRules = builder.NewFieldBuilder().
	MustAddLikeRule(regexutil.AlphabethRegexp(regexutil.DefaultDescriptionAlphabeth)).
	AddMaxLength(1024 * 16)

var langkaWordTupleCreateDataValidationRules = builder.NewBuilder().
	AddFieldBuilder("sourceWord", langkaWordValidationRules).
	AddFieldBuilder("destinationWords", builder.NewFieldBuilder().
		AddListOf(langkaWordValidationRules),
	).
	AddFieldBuilder("description", langkaWordTupleDescriptionValidationRules)

var LangkaWordSetCreateDataValidationRules = builder.NewBuilder().
	AddFieldBuilder("name", langkaWordSetNameValidationRules).
	AddFieldBuilder("fromLanguage", langkaLanguageNameValidationRules).
	AddFieldBuilder("toLanguage", langkaLanguageNameValidationRules).
	AddFieldBuilder("wordTuples", builder.
		NewFieldBuilder().
		AddRequired().
		AddListOfObjects(langkaWordTupleCreateDataValidationRules),
	).
	AddFieldBuilder("description", langkaWordSetDescriptionValidationRules).
	MustBuild()
