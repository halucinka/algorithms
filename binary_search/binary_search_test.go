package binary_search

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_BinarySearch(t *testing.T) {
	Convey("Test BinarySearch", t, func() {
		So(BinarySearch([]int{1, 2, 3}, 2), ShouldEqual, 1)
		So(BinarySearch([]int{1, 2, 3}, 47), ShouldEqual, -1)
		So(BinarySearch([]int{1, 2, 3}, 0), ShouldEqual, -1)
		So(BinarySearch([]int{1, 2, 4}, 3), ShouldEqual, -1)
		So(BinarySearch([]int{1, 2}, 2), ShouldEqual, 1)
		So(BinarySearch([]int{1, 2}, 1), ShouldEqual, 0)
		So(BinarySearch([]int{}, 47), ShouldEqual, -1)
	})
}
