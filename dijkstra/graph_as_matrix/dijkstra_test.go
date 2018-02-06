package graph_as_matrix

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetDistancesFromStart(t *testing.T) {
	Convey("Test GetDistancesFromStart", t, func() {
		graph := [][]int{
			{0, 1, 1},
			{1, 0, 0},
			{1, 0, 0},
		}
		So(GetDistancesFromStart(graph, 0), ShouldResemble, []int{0, 1, 1})
	})
	Convey("Test GetDistancesFromStart", t, func() {
		graph := [][]int{
			{0, 1, 1, 0},
			{1, 0, 1, 0},
			{1, 0, 0, 1},
			{0, 0, 1, 0},
		}
		So(GetDistancesFromStart(graph, 0), ShouldResemble, []int{0, 1, 1, 2})
	})
	Convey("Test GetDistancesFromStart - 2 components", t, func() {
		graph := [][]int{
			{0, 1, 0, 0},
			{1, 0, 0, 0},
			{0, 0, 0, 1},
			{0, 0, 1, 0},
		}
		So(GetDistancesFromStart(graph, 0), ShouldResemble, []int{0, 1, -1, -1})
	})
	Convey("Test GetDistancesFromStart - 4 components, no edges", t, func() {
		graph := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		So(GetDistancesFromStart(graph, 0), ShouldResemble, []int{0, -1, -1, -1})
	})
}