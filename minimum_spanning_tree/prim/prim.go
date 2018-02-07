package prim

import "container/heap"

type Edge struct {
	vertex1 int
	vertex2 int
	weight  int
}

type Edges struct {
	edges []*Edge
}

var Pointers = []int{}

type Vertex struct {
	id               int
	weight           int
	choosenNeighbour int
}

type Heap []*Vertex

func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	return h[i].weight < h[j].weight
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	Pointers[h[i].id], Pointers[h[j].id] = Pointers[h[j].id], Pointers[h[i].id]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*Vertex))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	Pointers[x.id] = -1
	return x
}

func Prim(graph [][]*Vertex) []*Edge {
	mst := []*Edge{}

	infinity := 100000000
	Pointers = make([]int, len(graph))
	h := Heap{}
	for i := 0; i < len(graph); i++ {
		Pointers[i] = i
		h = append(h, &Vertex{id: i, weight: infinity, choosenNeighbour: -1})
	}
	start := 0
	h[start] = &Vertex{id: start, weight: 0}
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

		delete(notUsedVertices, u.id)
		if u.id != start {
			mst = append(mst, &Edge{vertex1: u.id, vertex2: u.choosenNeighbour, weight: u.weight})
		}
		// for all neighbours
		for _, vertex := range graph[u.id] {
			neighbour := vertex.id
			value := vertex.weight
			if notUsedVertices[neighbour] && distances[neighbour] > value {
				distances[neighbour] = value
				h[Pointers[neighbour]] = &Vertex{id: neighbour, weight: value, choosenNeighbour: u.id}
				heap.Fix(&h, Pointers[neighbour])
			}

		}
	}
	return mst
}
