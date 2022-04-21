package regexutil

import (
	"regexp"
)

const ASCIILowercase = "abcdefghijklmnopqrstuvwxyz"
const ASCIIUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Digits = "0123456789"
const ASCIISpecial = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
const ASCIIWhite = "\n\t\v "

const PolishLowercase = "óąćęłńśźż"
const PolishUppercase = "ÓĄĆĘŁŃŚŹŻ"

const DefaultNameAlphabeth = ASCIILowercase + ASCIIUppercase + PolishLowercase + PolishUppercase + Digits
const ExtendedNameAlphabeth = DefaultNameAlphabeth + ASCIISpecial + " "

const DefaultDescriptionAlphabeth = ASCIILowercase + ASCIIUppercase + PolishLowercase + PolishUppercase + Digits + ASCIISpecial + ASCIIWhite

// Regex, which allows all characters from alphabeth provided
func AlphabethRegexp(alpha string) (res string) {
	res = "^["
	for _, r := range alpha {
		res += SanitizeCharacter(r)
	}
	res += "]*$"

	return res
}

func CompiledAlphabethRegexp(alpha string) *regexp.Regexp {
	return regexp.MustCompile(AlphabethRegexp(alpha))
}
