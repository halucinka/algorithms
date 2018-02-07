package graph_as_matrix

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var infinity = 100000000

func Test_GetDistancesFromStart(t *testing.T) {
	Convey("Test GetDistancesFromStart", t, func() {
		graph := [][]*Vertex{
			{{Id: 1, Distance: 1}, {Id: 2, Distance: 10}},
			{{Id: 0, Distance: 1}},
			{{Id: 0, Distance: 1}},
		}

		So(GetDistancesFromStart(graph, 0), ShouldResemble, []int{0, 1, 10})
	})
	Convey("Test GetDistancesFromStart", t, func() {
		graph := [][]*Vertex{
			{{Id: 1, Distance: 1}, {Id: 2, Distance: 10}},
			{{Id: 0, Distance: 1}, {Id: 2, Distance: 1}},
			{{Id: 2, Distance: 1}},
		}
		So(GetDistancesFromStart(graph, 0), ShouldResemble, []int{0, 1, 2})
	})
	Convey("Test GetDistancesFromStart - 2 components", t, func() {
		graph := [][]*Vertex{
			{{Id: 1, Distance: 1}},
			{{Id: 0, Distance: 1}},
			{{Id: 3, Distance: 1}},
			{{Id: 2, Distance: 1}},
		}
		So(GetDistancesFromStart(graph, 0), ShouldResemble, []int{0, 1, infinity, infinity})
	})
	Convey("Test GetDistancesFromStart - 4 components, no edges", t, func() {
		graph := [][]*Vertex{{}, {}, {}, {}}
		So(GetDistancesFromStart(graph, 0), ShouldResemble, []int{0, infinity, infinity, infinity})
	})
	Convey("Test GetDistancesFromStart - 2 components", t, func() {
		graph := [][]*Vertex{
			{{Id: 1, Distance: 1}}, {},
		}
		So(GetDistancesFromStart(graph, 0), ShouldResemble, []int{0, 1})
		So(GetDistancesFromStart(graph, 1), ShouldResemble, []int{infinity, 0})
	})
}
