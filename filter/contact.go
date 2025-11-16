package filter

import (
	"regexp"
	"strings"
)

func physicalAddress(text string) []string {
	pattern := regexp.MustCompile(`(?i)([\w\s.-]+?(?:Avenue|Avn|blvd|Street|str|Location))`)
	matches := pattern.FindAllString(text, -1)

	return matches
}

func Contact(text string) string {
	if text == "" {
		return text
	}

	newText := text
	addresses := physicalAddress(text)

	for _, address := range addresses {
		addr := strings.Trim(address, " ")
		newText = strings.ReplaceAll(newText, addr, "00A A0A0 AAA, AAAA, AA 0000")
	}

	return newText
}
