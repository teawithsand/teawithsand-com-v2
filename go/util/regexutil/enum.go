package regexutil

import (
	"regexp"
)

// Regex, which allows only set of values provided.
func EnumRegexp(values []string) (res string) {
	res += "^("
	for i, v := range values {
		res += SanitizeText(v)
		if i != len(values)-1 {
			res += "|"
		}
	}
	res += ")$"
	return res
}

func CompiledEnumRegexp(values []string) *regexp.Regexp {
	return regexp.MustCompile(EnumRegexp(values))
}
