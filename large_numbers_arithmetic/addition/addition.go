package addition

import (
	"strconv"
)

func Add(a, b string) string {
	result := ""
	maxLength := len(a)
	if len(b) > maxLength {
		maxLength = len(b)
	}
	carry := 0
	for k := 0; k < maxLength; k++ {
		addition := getKthDigitFromEnd(a, k) + getKthDigitFromEnd(b, k) + carry
		if addition > 9 {
			carry = 1
			addition = addition % 10
		}
		result = strconv.Itoa(addition) + result
	}
	if carry > 0 {
		result = strconv.Itoa(carry) + result
	}
	return result
}

func getDigit(s string, i int) int {
	digit, err := strconv.Atoi(string(s[i]))
	if err != nil {
		panic("Can't parse digit.")
	}
	return digit
}

func getKthDigitFromEnd(s string, k int) int {
	if k < len(s) {
		return getDigit(s, len(s)-1-k)
	}
	return 0
}
