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
		So(mst[0].weight, ShouldEqual, 1)
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
		So(mst[0].weight, ShouldEqual, 1)
	})
}
