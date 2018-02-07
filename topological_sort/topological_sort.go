package topological_sort

import "fmt"

func GetOrdering(graph [][]int) []int {
	N := len(graph)
	ordering := []int{}
	edgesIn := make([]int, N)
	queue := []int{}
	for i := 0; i < N; i++ {
		for _, neighbour := range graph[i] {
			edgesIn[neighbour]++
		}
	}
	for i := 0; i < N; i++ {
		if edgesIn[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		//pop
		u := queue[0]
		queue = queue[1:]
		ordering = append(ordering, u)
		for _, neighbour := range graph[u] {
			edgesIn[neighbour]--
			if edgesIn[neighbour] == 0 {
				queue = append(queue, neighbour)
			}
		}
	}

	if len(ordering) < N {
		fmt.Println(ordering)
		return []int{} // it's not a DAG, ordering not possible
	}
	return ordering
}
