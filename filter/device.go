package filter

import (
	"net"
)

func ipAddressCheck(token string) string {
	ip := net.ParseIP(token)
	if ip == nil {
		return token
	}

	return "XXX.XXX.XXX.XXX"
}

func macAddressCheck(token string) string {
	_, err := net.ParseMAC(token)
	if err != nil {
		return token
	}

	return "XX:XX:XX:XX:XX:XX"
}

func imeiCheck(token string) string {
	if len(token) != 15 {
		return token
	}

	if !IsAllDigitsASCII(token) {
		return token
	}

	return "000000000000000"
}

func InitDevice(text string) string {
	if text == "" {
		return text
	}

	newText := text
	tokens := Tokenize(newText)

	for _, token := range tokens {
		ip := ipAddressCheck(token)
		if valid, s := ReplaceFlaggedToken(token, ip, newText); valid {
			newText = s
			continue
		}

		mac := macAddressCheck(token)
		if valid, s := ReplaceFlaggedToken(token, mac, newText); valid {
			newText = s
			continue
		}

		imei := imeiCheck(token)
		if valid, s := ReplaceFlaggedToken(token, imei, newText); valid {
			newText = s
			continue
		}
	}

	return newText
}
