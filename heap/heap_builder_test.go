package heap

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHeapBuilder_HeapSort(t *testing.T) {
	Convey("Test HeapSort", t, func() {
		heapBuilder := NewHeapBuilder()
		So(heapBuilder.HeapSort([]int{2, 3, 1}), ShouldResemble, []int{1, 2, 3})
		So(heapBuilder.HeapSort([]int{2, 1}), ShouldResemble, []int{1, 2})
		So(heapBuilder.HeapSort([]int{1}), ShouldResemble, []int{1})
		So(heapBuilder.HeapSort([]int{}), ShouldResemble, []int{})
		So(heapBuilder.HeapSort(
			[]int{1, 16, 5, 30, 27, 17, 20, 2, 57, 3, 90}),
			ShouldResemble,
			[]int{1, 2, 3, 5, 16, 17, 20, 27, 30, 57, 90})

	})
}

func TestHeapBuilder_Insert(t *testing.T) {
	Convey("Test Insert", t, func() {
		heapBuilder := NewHeapBuilder()
		So(heapBuilder.Insert(4, []int{3, 2, 1}), ShouldResemble, []int{4, 3, 1, 2})
	})
}

func TestHeapBuilder_ExtractMaximum(t *testing.T) {
	Convey("Test ExtractMaximum", t, func() {
		heapBuilder := NewHeapBuilder()
		maximum, newHeap := heapBuilder.ExtractMaximum([]int{3, 2, 1})
		So(maximum, ShouldEqual, 3)
		So(newHeap, ShouldResemble, []int{2, 1, 3})
	})
}

func TestHeapBuilder_BuildHeap(t *testing.T) {
	Convey("Test BuildHeap", t, func() {
		heapBuilder := NewHeapBuilder()
		So(heapBuilder.BuildHeap([]int{1, 2, 3}), ShouldResemble, []int{3, 2, 1})
		So(heapBuilder.BuildHeap([]int{1}), ShouldResemble, []int{1})
		So(heapBuilder.BuildHeap([]int{}), ShouldResemble, []int{})
		So(heapBuilder.BuildHeap(
			[]int{1, 16, 5, 30, 27, 17, 20, 2, 57, 3, 90}),
			ShouldResemble,
			[]int{90, 57, 20, 30, 27, 17, 5, 2, 1, 3, 16})
	})
}

func TestHeapBuilder_Heapify(t *testing.T) {
	Convey("Test BubbleDown", t, func() {
		heapBuilder := heapBuilder{}
		So(heapBuilder.BubbleDown(0, []int{1, 2, 3}, 3), ShouldResemble, []int{3, 2, 1})
		So(heapBuilder.BubbleDown(0, []int{1, 2}, 2), ShouldResemble, []int{2, 1})
		So(heapBuilder.BubbleDown(0, []int{2, 1}, 2), ShouldResemble, []int{2, 1})
	})
}

func TestHeapBuilder(t *testing.T) {
	Convey("Test", t, func() {
		heapBuilder := NewHeapBuilder()
		array := []int{1, 2, 3}
		array = heapBuilder.Insert(4, array)
		array = heapBuilder.Insert(5, array)
		array = heapBuilder.Insert(10, array)
		maximum, array := heapBuilder.ExtractMaximum(array)
		So(maximum, ShouldEqual, 10)
		So(array, ShouldResemble, []int{5, 4, 3, 2, 1, 10})
	})
}
