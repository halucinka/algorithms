package lomuto_quick_sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_LomutoQuickSort(t *testing.T) {
	Convey("Test LomutoQuickSort", t, func() {
		So(LomutoQuickSort([]int{4, 3, 2, 1}), ShouldResemble, []int{1, 2, 3, 4})
		So(LomutoQuickSort([]int{}), ShouldResemble, []int{})
		So(LomutoQuickSort([]int{2}), ShouldResemble, []int{2})
		So(LomutoQuickSort([]int{2, 1}), ShouldResemble, []int{1, 2})
		So(LomutoQuickSort([]int{3, 2, 1}), ShouldResemble, []int{1, 2, 3})
		So(LomutoQuickSort([]int{1, 2, 3}), ShouldResemble, []int{1, 2, 3})
		So(LomutoQuickSort([]int{1, 2, 1, 1, 1}), ShouldResemble, []int{1, 1, 1, 1, 2})
	})
}
