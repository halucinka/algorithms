package naive

func MatchString(text, pattern string) (bool, int) {
	for i := 0; i < len(text)-len(pattern)+1; i++ {
		if isMatch(pattern, text[i:i+len(pattern)]) {
			return true, i
		}
	}
	return false, -1
}
func isMatch(pattern string, text string) bool {
	for j := 0; j < len(pattern); j++ {
		if pattern[j] != text[j] {
			return false
		}
	}
	return true
}
