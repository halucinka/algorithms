package optimal_solution

import (
	"strconv"
)

func Multiply(a, b string) string {
	M := len(a)
	N := len(b)
	result := make([]int, M+N)

	for k := 0; k < M; k++ {
		for l := 0; l < N; l++ {
			result[k+l] += getKthDigitFromEnd(a, k) * getKthDigitFromEnd(b, l)
		}
	}
	result = postprocess(result)
	return transormResultToString(result)
}

func transormResultToString(result []int) string {
	resultString := ""
	nonZeroSeen := false
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != 0 || nonZeroSeen {
			resultString += strconv.Itoa(result[i])
			nonZeroSeen = true
		}
	}
	if len(resultString) == 0 {
		return "0"
	}
	return resultString
}

func postprocess(result []int) []int {
	var carry int
	for i := 0; i < len(result)-1; i++ {
		carry = result[i] / 10
		result[i] %= 10
		result[i+1] += carry
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
