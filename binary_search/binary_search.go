package binary_search

func BinarySearch(array []int, element int) int {
	left := 0
	right := len(array) - 1

	for left >= 0 && right < len(array) && left <= right {
		middle := (right-left)/2 + left
		if element < array[middle] {
			right = middle - 1
		} else if element > array[middle] {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
