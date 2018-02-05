package quick_sort

import "math/rand"

func QuickSort(array []int) []int {
	return quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, left, right int) []int {

	if left >= right {
		return array
	}

	array, middle := randomizedPartitioningOfSubarray(array, left, right)
	array = quickSort(array, left, middle-1)
	array = quickSort(array, middle+1, right)

	return array
}

func randomizedPartitioningOfSubarray(array []int, left, right int) ([]int, int) {
	pivot := left + rand.Intn(right-left)
	array = swap(array, right, pivot)
	return paritioningOfSubarray(array, right, left)
}

func paritioningOfSubarray(array []int, right int, left int) ([]int, int) {
	pivotValue := array[right]
	i := left - 1
	for j := left; j < right; j++ {
		if array[j] <= pivotValue {
			i++
			array = swap(array, i, j)
		}
	}
	array = swap(array, i+1, right)
	return array, i + 1
}

func swap(array []int, a, b int) []int {
	tmp := array[a]
	array[a] = array[b]
	array[b] = tmp
	return array
}
