package graph_as_matrix

import (
	"container/heap"
)

var Pointers = []int{}

type Vertex struct {
	Id       int
	Distance int
}

type Heap []*Vertex

func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	return h[i].Distance < h[j].Distance
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	Pointers[h[i].Id], Pointers[h[j].Id] = Pointers[h[j].Id], Pointers[h[i].Id]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*Vertex))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	Pointers[x.Id] = -1
	return x
}

func GetDistancesFromStart(graph [][]*Vertex, start int) []int {
	infinity := 100000000
	Pointers = make([]int, len(graph))
	h := Heap{}
	for i := 0; i < len(graph); i++ {
		Pointers[i] = i
		h = append(h, &Vertex{Id: i, Distance: infinity})
	}
	h[start] = &Vertex{Id: start, Distance: 0}
	heap.Init(&h)

	notUsedVertices := map[int]bool{}
	for i := range graph {
		notUsedVertices[i] = true
	}
	N := len(graph)
	distances := make([]int, N)
	for index := range distances {
		distances[index] = infinity
	}
	distances[start] = 0
	for len(notUsedVertices) > 0 {
		//extractMin
		u := heap.Pop(&h).(*Vertex)

		delete(notUsedVertices, u.Id)

		// for all neighbours
		for _, vertex := range graph[u.Id] {
			neighbour := vertex.Id
			value := vertex.Distance
			if distances[neighbour] > distances[u.Id]+value {
				distances[neighbour] = distances[u.Id] + value
				h[Pointers[neighbour]] = &Vertex{neighbour, distances[u.Id] + value}
				heap.Fix(&h, Pointers[neighbour])
			}

		}
	}
	return distances
}
