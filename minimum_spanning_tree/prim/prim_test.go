package prim

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Prim(t *testing.T) {
	Convey("Test Prim", t, func() {
		graph := [][]*Vertex{
			{{id: 1, weight: 1}, {id: 2, weight: 1}},
			{{id: 0, weight: 1}},
			{{id: 0, weight: 1}},
		}

		mst := Prim(graph)

		So(mst, ShouldHaveLength, 2)
		So(mst[0].weight, ShouldEqual, 1)
		So(mst[1].weight, ShouldEqual, 1)
	})

	Convey("Test Prim", t, func() {
		graph := [][]*Vertex{
			{{id: 1, weight: 1}, {id: 2, weight: 2}},
			{{id: 0, weight: 1}, {id: 2, weight: 1}},
			{{id: 0, weight: 2}, {id: 1, weight: 1}},
		}

		mst := Prim(graph)

		So(mst, ShouldHaveLength, 2)
		So(mst[0].weight, ShouldEqual, 1)
		So(mst[1].weight, ShouldEqual, 1)
	})

	Convey("Test Prim", t, func() {
		graph := [][]*Vertex{
			{{id: 1, weight: 3}, {id: 4, weight: 1}},
			{{id: 0, weight: 3}, {id: 2, weight: 5}, {id: 4, weight: 4}},
			{{id: 1, weight: 5}, {id: 3, weight: 2}, {id: 4, weight: 6}},
			{{id: 2, weight: 2}, {id: 4, weight: 7}},
			{{id: 0, weight: 1}, {id: 1, weight: 4}, {id: 2, weight: 6}, {id: 3, weight: 7}},
		}

		mst := Prim(graph)

		So(mst, ShouldHaveLength, 4)
		So(mst[0].weight, ShouldEqual, 1)
		So(mst[1].weight, ShouldEqual, 3)
		So(mst[2].weight, ShouldEqual, 5)
		So(mst[3].weight, ShouldEqual, 2)
	})
}
