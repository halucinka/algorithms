package multiplication

import (
	"fmt"
	"strconv"
)

func Multiply(a, b string) string {
	intermediateProducts := declare2DSlice(len(a), len(b))
	if a == "0" || b == "0" {
		return "0"
	}
	for k := 0; k < len(a); k++ {
		for l := 0; l < len(b); l++ {
			intermediateProducts[k][l] = getKthDigitFromEnd(a, k) * getKthDigitFromEnd(b, l)
		}
	}
	return postprocess(intermediateProducts, len(a), len(b))
}
func postprocess(products [][]int, N, M int) string {
	result := ""
	var carry int
	for i := 0; i < N+M-1; i++ {
		digit := carry
		carry = 0
		for k := max(0, i-M+1); k <= min(i, N-1); k++ {
			l := i - k

			digit += products[k][l]
		}
		if digit > 9 {
			carry = digit / 10
			digit = digit % 10
		}
		result = strconv.Itoa(digit) + result
	}
	if carry > 0 {
		result = strconv.Itoa(carry) + result
	}
	return result
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
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

func declare2DSlice(M int, N int) [][]int {
	matrix := make([][]int, M)
	for i := 0; i < M; i++ {
		matrix[i] = make([]int, N)
	}
	return matrix
}

// checking boundaries of double loop
func checkDiagonalTraversal() {
	products := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	M := 4
	N := 2
	for i := 0; i < N+M-1; i++ {
		fmt.Println("i", i)
		for k := max(0, i-M+1); k <= min(i, N-1); k++ {
			l := i - k
			fmt.Println("k, l", k, l)
			fmt.Println(products[k][l])
		}
		fmt.Println("-------")
	}

}
