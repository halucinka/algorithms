package topological_sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetOrdering(t *testing.T) {
	Convey("Test GetOrdering", t, func() {
		graph := [][]int{
			{1},
			{0},
		}
		So(GetOrdering(graph), ShouldResemble, []int{})
	})
	Convey("Test GetOrdering", t, func() {
		graph := [][]int{
			{2},
			{0},
			{1},
		}
		So(GetOrdering(graph), ShouldResemble, []int{})
	})
	Convey("Test GetOrdering", t, func() {
		graph := [][]int{
			{},
			{0},
			{1},
			{2},
		}
		So(GetOrdering(graph), ShouldResemble, []int{3, 2, 1, 0})
	})
	Convey("Test GetOrdering - 2 components", t, func() {
		graph := [][]int{
			{1},
			{},
			{},
			{2},
		}
		So(GetOrdering(graph), ShouldResemble, []int{0, 3, 1, 2})
	})
	Convey("Test GetDistancesFromStart - 4 components, no edges", t, func() {
		graph := [][]int{
			{},
			{},
			{},
			{},
		}
		So(GetOrdering(graph), ShouldResemble, []int{0, 1, 2, 3})
	})
	Convey("Test GetDistancesFromStart - 2 components", t, func() {
		graph := [][]int{
			{1},
			{},
		}
		So(GetOrdering(graph), ShouldResemble, []int{0, 1})
	})
}
