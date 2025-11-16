package filter

import (
	"regexp"
)

func physicalAddress(token string) string {
	var pattern = regexp.MustCompile(`^\d{1,5}[A-Za-z]?\s[A-Za-z0-9\s,\-#]+$`)

	if pattern.MatchString(token) {
		return "00A A0A0 AAA, AAAA, AA 0000"
	}

	return token
}

func Contact(text string) string {
	if text == "" {
		return text
	}

	newText := text
	tokens := Tokenize(newText)

	for _, token := range tokens {
		address := physicalAddress(token)
		if valid, s := ReplaceFlaggedToken(token, address, newText); valid {
			newText = s
			continue
		}
	}

	return newText
}
