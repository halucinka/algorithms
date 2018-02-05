package k_th_smallest_element

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetKthSmallestElement(t *testing.T) {
	Convey("Test GetKthSmallestElement", t, func() {
		So(GetKthSmallestElement([]int{}, 1), ShouldEqual, -1)
		So(GetKthSmallestElement([]int{1, 2, 3}, 1), ShouldEqual, 1)
		So(GetKthSmallestElement([]int{2, 4, 3, 1, 6}, 3), ShouldEqual, 3)
	})
}
