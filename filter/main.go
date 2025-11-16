package filter

import "strings"

func IsAllDigitsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}

	return true
}

func Tokenize(text string) []string {
	rawTokens := strings.Fields(text)

	var tokens []string
	signs := ",.?!():;\"'"

	for _, token := range rawTokens {
		if token != "" {
			tokens = append(tokens, strings.Trim(token, signs))
		}
	}

	return tokens
}

func ReplaceFlaggedToken(token1 string, token2 string, text string) (bool, string) {
	if token1 == token2 {
		return false, text
	}

	return true, strings.ReplaceAll(text, token1, token2)
}
