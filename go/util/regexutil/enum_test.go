package regexutil_test

import (
	"testing"

	"github.com/teawithsand/webpage/util/regexutil"
)

func TestEnumRegexp(t *testing.T) {
	re := regexutil.CompiledEnumRegexp([]string{
		"pl",
		"en",
		"fr",
	})

	do := func(ok bool, text string) {
		if t.Failed() {
			return
		}

		if re.MatchString(text) != ok {
			t.Error("expected", ok, "got", !ok, "for", text)
		}
	}

	do(true, "pl")
	do(false, "plpl")
	do(true, "en")
	do(false, "plen")
	do(false, "")
}
