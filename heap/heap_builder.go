package heap

import "math"

type HeapBuilder interface {
	BuildHeap(array []int) []int
	Insert(element int, array []int) []int
	HeapSort(array []int) []int
	ExtractMaximum(array []int) (int, []int)
}

func NewHeapBuilder() HeapBuilder {
	return &heapBuilder{}
}

type heapBuilder struct {
}

func (this heapBuilder) HeapSort(array []int) []int {
	array = this.BuildHeap(array)
	for heapSize := len(array) - 1; heapSize > 0; heapSize-- {
		array = this.Swap(0, heapSize, array)
		array = this.Heapify(0, array, heapSize)
	}
	return array
}

func (this heapBuilder) Insert(element int, array []int) []int {
	array = append(array, element)
	i := len(array) - 1
	for i > 0 && array[i] > this.parent(i) {
		array = this.Swap(i, this.parent(i), array)
		i = this.parent(i)
	}
	return array
}

func (this heapBuilder) BuildHeap(array []int) []int {
	for i := len(array) / 2; i >= 0; i-- {
		array = this.Heapify(i, array, len(array))
	}
	return array
}

func (this heapBuilder) ExtractMaximum(array []int) (int, []int) {
	heapSize := len(array)
	if len(array) < 1 {
		return 0, []int{}
	}
	maximum := array[0]
	array = this.Swap(0, heapSize-1, array)
	array = this.Heapify(0, array, heapSize-1)
	return maximum, array
}

func (this heapBuilder) Swap(a, b int, array []int) []int {
	tmp := array[a]
	array[a] = array[b]
	array[b] = tmp
	return array
}

func (this heapBuilder) Heapify(i int, array []int, heapSize int) []int {
	if heapSize < 2*(i+1)-1 {
		return array
	}

	childForSwap := this.getChildForSwap(array, i, heapSize)

	if childForSwap == -1 {
		return array
	}

	array = this.Swap(childForSwap, i, array)
	return this.Heapify(childForSwap, array, heapSize)
}

func (heapBuilder) getChildForSwap(array []int, i, heapSize int) int {
	leftChildIndex := 2*i + 1
	rightChildIndex := 2*i + 2

	if rightChildIndex < heapSize &&
		array[i] < array[rightChildIndex] &&
		array[rightChildIndex] > array[leftChildIndex] {
		return rightChildIndex
	}
	if leftChildIndex < heapSize && array[i] < array[leftChildIndex] {
		return leftChildIndex
	}

	return -1
}
func (heapBuilder) parent(i int) int {
	return int(math.Floor(float64((i - 1) / 2)))
}
