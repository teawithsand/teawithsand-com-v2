package regexutil

import "unicode"

var replacementMap = map[rune]string{
	'\n': "\\n",
	'\v': "\\v",
	'\t': "\\t",
	'\r': "\\r",
}

func SanitizeCharacter(r rune) (res string) {
	rep, ok := replacementMap[r]

	if ok {
		res += rep
	} else if r > unicode.MaxASCII {
		if unicode.IsLetter(r) {
			res += string(r)
		} else {
			res += "\\" // just comment out any char that may be there
			res += string(r)
		}
	} else {
		res += string(r)
	}
	return
}

func SanitizeText(text string) (res string) {
	for _, r := range text {
		res += SanitizeCharacter(r)
	}
	return
}
