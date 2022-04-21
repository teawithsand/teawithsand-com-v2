package regexutil

import "regexp"

const TrimmedRegexp = "^\\S.*\\S$"

var CompiledTrimmedRegexp = regexp.MustCompile(TrimmedRegexp)
