package quick_sort

import "math/rand"

func QuickSort(array []int) []int {

	if len(array) <= 1 {
		arrayCopy := make([]int, len(array))
		copy(arrayCopy, array)
		return arrayCopy
	}

	pivot := array[rand.Intn(len(array))]

	var less, middle, more []int

	for _, value := range array {
		switch {
		case value < pivot:
			less = append(less, value)
		case value == pivot:
			middle = append(middle, value)
		case value > pivot:
			more = append(more, value)
		}
	}

	less = QuickSort(less)
	more = QuickSort(more)

	less = append(less, middle...)
	less = append(less, more...)
	return less
}
