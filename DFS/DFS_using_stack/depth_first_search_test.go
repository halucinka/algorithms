package DFS_using_stack

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetNumberOfComponents(t *testing.T) {

	Convey("Test GetNumberOfComponents - one component", t, func() {
		graph := [][]int{
			{1, 2},
			{0, 2},
			{0, 1, 3},
			{2},
		}
		So(GetNumberOfComponents(graph), ShouldEqual, 1)
	})
	Convey("Test GetNumberOfComponents - 2 components, isolated vertex", t, func() {
		graph := [][]int{
			{1, 2},
			{0, 2},
			{0, 1},
			{},
		}
		So(GetNumberOfComponents(graph), ShouldEqual, 2)
	})
	Convey("Test GetNumberOfComponents - 2 components", t, func() {
		graph := [][]int{
			{1},
			{0},
			{3},
			{2},
		}
		So(GetNumberOfComponents(graph), ShouldEqual, 2)
	})
	Convey("Test GetNumberOfComponents - 4 components, no edges", t, func() {
		graph := [][]int{
			{},
			{},
			{},
			{},
		}
		So(GetNumberOfComponents(graph), ShouldEqual, 4)
	})
}
