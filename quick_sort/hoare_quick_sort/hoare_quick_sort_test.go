package hoare_quick_sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_HoareQuickSort(t *testing.T) {
	Convey("Test HoareQuickSort", t, func() {
		So(HoareQuickSort([]int{4, 3, 2, 1}), ShouldResemble, []int{1, 2, 3, 4})
		So(HoareQuickSort([]int{}), ShouldResemble, []int{})
		So(HoareQuickSort([]int{2}), ShouldResemble, []int{2})
		So(HoareQuickSort([]int{2, 1}), ShouldResemble, []int{1, 2})
		So(HoareQuickSort([]int{3, 2, 1}), ShouldResemble, []int{1, 2, 3})
		So(HoareQuickSort([]int{1, 2, 3}), ShouldResemble, []int{1, 2, 3})
		So(HoareQuickSort([]int{1, 2, 1, 1, 1}), ShouldResemble, []int{1, 1, 1, 1, 2})
	})
}
