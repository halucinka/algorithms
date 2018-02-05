package rabin_karp

import "math"

func MatchString(text, pattern string) (bool, int) {
	if len(text) < len(pattern) {
		return false, -1
	}
	return matchString(text, pattern, 52, 47)
}

func matchString(text, pattern string, d, q int) (bool, int) {
	textLength := len(text)
	patternLength := len(pattern)
	h := int(math.Pow(float64(d), float64((patternLength-1)))) % q
	patternValue, textValue := preprocess(d, q, pattern, text)

	for position := 0; position < textLength-patternLength; position++ {
		if patternValue == textValue {
			if isMatch(pattern, text[position:position+patternLength]) {
				return true, position
			}
		}
		if position < textLength-patternLength-1 {
			textValue = (d*(textValue-int(text[position])*h) + int(text[position+patternLength])) % q
		}
	}
	return false, -1
}

func preprocess(d, q int, pattern, text string) (patternValue, textValue int) {
	patternValue = 0
	textValue = 0
	for i := 0; i < len(pattern); i++ {
		patternValue = (d*patternValue + int(pattern[i])) % q
		textValue = (d*textValue + int(text[i])) % q
	}
	return patternValue, textValue
}

func isMatch(pattern string, text string) bool {
	for j := 0; j < len(pattern); j++ {
		if pattern[j] != text[j] {
			return false
		}
	}
	return true
}
