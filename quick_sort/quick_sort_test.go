package quick_sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_QuickSort(t *testing.T) {
	Convey("Test QuickSort", t, func() {
		So(QuickSort([]int{4, 3, 2, 1}), ShouldResemble, []int{1, 2, 3, 4})
		So(QuickSort([]int{}), ShouldResemble, []int{})
		So(QuickSort([]int{2}), ShouldResemble, []int{2})
		So(QuickSort([]int{1, 2, 3}), ShouldResemble, []int{1, 2, 3})
		So(QuickSort([]int{1, 2, 1, 1, 1}), ShouldResemble, []int{1, 1, 1, 1, 2})
	})
}
