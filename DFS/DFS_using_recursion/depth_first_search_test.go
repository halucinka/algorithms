package DFS_using_recursion

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetNumberOfComponents_recursive(t *testing.T) {

	Convey("Test GetNumberOfComponents_recursive - one component", t, func() {
		graph := [][]int{
			{1, 2},
			{0, 2},
			{0, 1, 3},
			{2},
		}
		So(GetNumberOfComponents_recursive(graph), ShouldEqual, 1)
	})
	Convey("Test GetNumberOfComponents_recursive - 2 components, isolated vertex", t, func() {
		graph := [][]int{
			{1, 2},
			{0, 2},
			{0, 1},
			{},
		}
		So(GetNumberOfComponents_recursive(graph), ShouldEqual, 2)
	})
	Convey("Test GetNumberOfComponents_recursive - 2 components", t, func() {
		graph := [][]int{
			{1},
			{0},
			{3},
			{2},
		}
		So(GetNumberOfComponents_recursive(graph), ShouldEqual, 2)
	})
	Convey("Test GetNumberOfComponents_recursive - 4 components, no edges", t, func() {
		graph := [][]int{
			{},
			{},
			{},
			{},
		}
		So(GetNumberOfComponents_recursive(graph), ShouldEqual, 4)
	})
}
