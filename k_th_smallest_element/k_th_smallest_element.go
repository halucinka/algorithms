package k_th_smallest_element

import "math/rand"

func GetKthSmallestElement(array []int, k int) int {
	if len(array) == 0 {
		return -1
	}
	return getKthSmallestElement(array, k, 0, len(array)-1)
}

func getKthSmallestElement(array []int, k, left, right int) int {
	if left == right {
		return array[left]
	}
	array, middle := randomizedPartitioningOfSubarray(array, left, right)
	lengthOfLeftSubArray := middle - left + 1
	if lengthOfLeftSubArray > k {
		getKthSmallestElement(array, left, middle-1, k)
	} else if lengthOfLeftSubArray < k {
		getKthSmallestElement(array, middle+1, right, k-lengthOfLeftSubArray)
	} else {
		return array[middle]
	}
	return -1
}

func randomizedPartitioningOfSubarray(array []int, left, right int) ([]int, int) {
	pivot := left + rand.Intn(right-left) - 1
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
