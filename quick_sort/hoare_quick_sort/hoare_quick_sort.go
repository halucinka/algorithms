package hoare_quick_sort

import "math/rand"

func HoareQuickSort(array []int) []int {
	return quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, left, right int) []int {

	if left >= right {
		return array
	}

	array, middle := paritioningOfSubarray(array, left, right)
	array = quickSort(array, left, middle)
	array = quickSort(array, middle+1, right)

	return array
}

func paritioningOfSubarray(array []int, left, right int) ([]int, int) {
	pivot := left + rand.Intn(right-left)
	pivotValue := array[pivot]
	l := left
	r := right
	for l < r {
		for l <= right && array[l] < pivotValue {
			l++
		}
		for r >= left && array[r] > pivotValue {
			r--
		}
		if l < r {
			array = swap(array, l, r)
			l++
			r--
		}
	}

	return array, r
}

func swap(array []int, a, b int) []int {
	tmp := array[a]
	array[a] = array[b]
	array[b] = tmp
	return array
}
