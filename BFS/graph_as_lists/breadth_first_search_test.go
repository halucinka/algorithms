package graph_as_lists

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetDistancesFromStart(t *testing.T) {
	Convey("Test GetDistancesFromStart", t, func() {
		graph := [][]int{
			{1, 2},
			{0},
			{0},
		}
		So(GetDistancesFromStart(graph, 0), ShouldResemble, []int{0, 1, 1})
	})
	Convey("Test GetDistancesFromStart", t, func() {
		graph := [][]int{
			{1, 2},
			{0, 2},
			{0, 3},
			{2},
		}
		So(GetDistancesFromStart(graph, 0), ShouldResemble, []int{0, 1, 1, 2})
	})
	Convey("Test GetDistancesFromStart - 2 components", t, func() {
		graph := [][]int{
			{1},
			{0},
			{3},
			{2},
		}
		So(GetDistancesFromStart(graph, 0), ShouldResemble, []int{0, 1, -1, -1})
	})
	Convey("Test GetDistancesFromStart - 4 components, no edges", t, func() {
		graph := [][]int{
			{},
			{},
			{},
			{},
		}
		So(GetDistancesFromStart(graph, 0), ShouldResemble, []int{0, -1, -1, -1})
	})
}

func Test_GetNumberOfComponents(t *testing.T) {

	Convey("Test GetNumberOfComponents - one component", t, func() {
		graph := [][]int{
			{1, 2},
			{0, 2},
			{0, 3},
			{2},
		}
		So(GetNumberOfComponents(graph), ShouldEqual, 1)
	})
	Convey("Test GetNumberOfComponents - 2 components, isolated vertex", t, func() {
		graph := [][]int{
			{1, 2},
			{0, 2},
			{0},
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
