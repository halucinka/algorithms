package kruskal

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Kruskal(t *testing.T) {
	Convey("Test Kruskal", t, func() {
		graph := [][]int{
			{0, 1, 1},
			{1, 0, 0},
			{1, 0, 0},
		}
		mst := Kruskal(graph)
		So(mst, ShouldHaveLength, 2)
		So(mst[0].weight, ShouldEqual, 1)
		So(mst[1].weight, ShouldEqual, 1)
	})
	Convey("Test Kruskal", t, func() {
		graph := [][]int{
			{0, 1, 2},
			{1, 0, 1},
			{2, 1, 0},
		}
		mst := Kruskal(graph)
		So(mst, ShouldHaveLength, 2)
		So(mst[0].weight, ShouldEqual, 1)
		So(mst[1].weight, ShouldEqual, 1)
	})
	Convey("Test Kruskal", t, func() {
		graph := [][]int{
			{0, 3, 0, 0, 1},
			{3, 0, 5, 0, 4},
			{0, 5, 0, 2, 6},
			{0, 0, 2, 0, 7},
			{1, 4, 6, 7, 0},
		}
		mst := Kruskal(graph)
		So(mst, ShouldHaveLength, 4)
		So(mst[0].weight, ShouldEqual, 1)
		So(mst[1].weight, ShouldEqual, 2)
		So(mst[2].weight, ShouldEqual, 3)
		So(mst[3].weight, ShouldEqual, 5)
	})
}
