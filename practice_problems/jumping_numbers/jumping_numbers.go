package jumping_numbers

// https://practice.geeksforgeeks.org/problems/jumping-numbers/0#ExpectOP

func GetJumpingNumbers(number int) []int {
	result := []int{0}
	for i := 1; i < min(10, number+1); i++ {
		stack := []int{i}
		result = append(result, i)
		for len(stack) > 0 {
			//pop
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			lastDigit := a % 10
			if lastDigit == 0 {
				newNumber := a*10 + 1
				if newNumber <= number {
					stack = append(stack, newNumber)
					result = append(result, newNumber)
				}
			} else if lastDigit == 9 {
				newNumber := a*10 + 8
				if newNumber <= number {
					stack = append(stack, newNumber)
					result = append(result, newNumber)
				}
			} else {
				newNumber1 := a*10 + a%10 - 1
				newNumber2 := a*10 + a%10 + 1
				if newNumber2 <= number {
					stack = append(stack, newNumber2)
				}
				if newNumber1 <= number {
					stack = append(stack, newNumber1)
					result = append(result, newNumber1)
				}
				if newNumber2 <= number {
					result = append(result, newNumber2)
				}
			}
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
