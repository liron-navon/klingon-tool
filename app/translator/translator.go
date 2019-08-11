package translator

import (
	"errors"
	"strings"
)

// TranslateEnglishToClingon - translates a text in english to a text in klingon unicode
func TranslateEnglishToKlingon(s string) (string, error) {
	var results []string

	for len(s) > 0 {
		var token string

		// we only have few multi character klingon symbols
		if strings.HasPrefix(s, "tlh") {
			token = s[0:3]
		} else if strings.HasPrefix(s, "ch") ||
			strings.HasPrefix(s, "ng") {
			token = s[0:2]
		} else {
			token = s[0:1]

			switch token {
			case "s", "d", "h":
				token = strings.ToUpper(token)
			case "S", "D", "Q", "H":
				// we don't want to mutate these letters
				break
			default:
				token = strings.ToLower(token)
			}
		}

		s = s[len(token):]
		translation := klingonDict[token]
		if translation == "" {
			return "", errors.New("unable to translate this text to klingon")
		}

		results = append(results, translation)
	}
	return strings.Join(results, " "), nil
}

// klingonDict - the klingon dict is our reference to translate tokens into klingon unicode symbols
var klingonDict = map[string]string{
	"a":   "0xF8D0",
	"b":   "0xF8D1",
	"ch":  "0xF8D2",
	"D":   "0xF8D3",
	"e":   "0xF8D4",
	"gh":  "0xF8D5",
	"H":   "0xF8D6",
	"i":   "0xF8D7",
	"j":   "0xF8D8",
	"l":   "0xF8D9",
	"m":   "0xF8DA",
	"n":   "0xF8DB",
	"ng":  "0xF8DC",
	"o":   "0xF8DD",
	"p":   "0xF8DE",
	"q":   "0xF8DF",
	"Q":   "0xF8E0",
	"r":   "0xF8E1",
	"S":   "0xF8E2",
	"t":   "0xF8E3",
	"tlh": "0xF8E4",
	"u":   "0xF8E5",
	"v":   "0xF8E6",
	"w":   "0xF8E7",
	"y":   "0xF8E8",
	"’":   "0xF8E9",
	" ":   "0x0020",
}
