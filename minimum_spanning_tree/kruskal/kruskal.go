package kruskal

import (
	"sort"
)

func (e Edges) Len() int {
	return len(e.edges)
}

func (e Edges) Less(i, j int) bool {
	return e.edges[i].weight < e.edges[j].weight
}
func (e Edges) Swap(i, j int) {
	e.edges[i], e.edges[j] = e.edges[j], e.edges[i]
}

var parents = []int{}
var ranks = []int{}

type Edge struct {
	vertex1 int
	vertex2 int
	weight  int
}

type Edges struct {
	edges []*Edge
}

func Kruskal(graph [][]int) []*Edge {
	mst := []*Edge{}
	edges := Edges{}
	parents = make([]int, len(graph))
	ranks = make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		parents[i] = i
	}
	for u := 0; u < len(graph); u++ {
		for v := u + 1; v < len(graph); v++ {
			if graph[u][v] != 0 {
				edges.edges = append(edges.edges, &Edge{vertex1: u, vertex2: v, weight: graph[u][v]})
			}
		}
	}
	sort.Sort(edges)

	for _, edge := range edges.edges {
		if find(edge.vertex1) != find(edge.vertex2) {
			mst = append(mst, edge)
			union(edge.vertex1, edge.vertex2)
		}
	}
	return mst
}

func union(u, v int) {
	rootU := find(u)
	rootV := find(v)

	if rootU == rootV {
		return
	}
	if ranks[rootU] < ranks[rootV] {
		parents[rootU] = rootV
	} else if ranks[rootU] > ranks[rootV] {
		parents[rootV] = rootU
	} else {
		parents[rootV] = rootU
		ranks[rootU]++
	}
	return
}

func find(v int) int {
	u := v
	for parents[u] != u {
		u = parents[u]
	}
	rootV := u
	u = v
	for parents[u] != u {
		parent := parents[u]
		parents[u] = rootV
		u = parent
	}
	return rootV
}
