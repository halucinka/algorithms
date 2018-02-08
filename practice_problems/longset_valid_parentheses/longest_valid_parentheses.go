package longset_valid_parentheses

// https://practice.geeksforgeeks.org/problems/longest-valid-parentheses/0

func GetLongestSubString(parentheses string) (int, int) {
	stack := []int{}
	stackStart := 0
	// lengthOfLastGoodString allows us to track previous good expressions. e.g. ()(), that means ()<-empty stack->()
	start := 0
	end := 0
	maxLength := -1
	for i := 0; i < len(parentheses); i++ {
		if rune(parentheses[i]) == 40 {
			// push: saving index of unresolved opening parenthesis
			stack = append(stack, i)
		} else {
			if len(stack) > 0 {
				//pop
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					indexOfPreviousOpen := stack[len(stack)-1]
					if maxLength < i-indexOfPreviousOpen {
						maxLength = i - indexOfPreviousOpen
						start = indexOfPreviousOpen + 1
						end = i
					}
				} else {
					// we found a good string, we can concatenate it with previous string we found
					// when stack was empty the last time
					length := i - stackStart + 1

					if maxLength < length {
						maxLength = length
						start = stackStart
						end = i
					}
				}
			} else {
				// not enough (, not correct expression, we have to start again.
				stackStart = i + 1
			}
		}
	}
	return start, end
}
