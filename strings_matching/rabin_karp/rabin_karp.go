package rabin_karp

import "math"

func MatchString(text, pattern string) (bool, int) {
	return matchString(text, pattern, 52, 47)
}

func matchString(text, pattern string, d, q int) (bool, int) {
	textLength := len(text)
	patternLength := len(pattern)

	if len(text) < patternLength {
		return false, -1
	}
	h := int(math.Pow(float64(d), float64((patternLength-1)))) % q
	patternValue, textValue, h := preprocess(d, q, pattern, text)

	for position := 0; position < textLength-patternLength+1; position++ {
		if patternValue == textValue {
			if isMatch(pattern, text[position:position+patternLength]) {
				return true, position
			}
		}
		if position < textLength-patternLength {
			textValue = (d*textValue - int(text[position])*h + int(text[position+patternLength])) % q
			textValue = (textValue + q) % q
		}
	}
	return false, -1
}

func preprocess(d, q int, pattern, text string) (patternValue, textValue, h int) {
	patternValue = 0
	textValue = 0
	h = 1
	for i := 0; i < len(pattern); i++ {
		patternValue = (d*patternValue + int(pattern[i])) % q
		textValue = (d*textValue + int(text[i])) % q
		h = (h * d) % q
	}
	return patternValue, textValue, h
}

func isMatch(pattern string, text string) bool {
	for j := 0; j < len(pattern); j++ {
		if pattern[j] != text[j] {
			return false
		}
	}
	return true
}
