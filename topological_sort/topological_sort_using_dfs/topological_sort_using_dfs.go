package topological_sort_using_dfs

var topologicalOrdering = []int{}
var colors = []int{}

func GetOrdering(graph [][]int) []int {
	ordering := DFS_recursive(graph)
	if len(ordering) < len(graph) {
		return []int{} // ordering not possible
	}
	return ordering
}

func DFS_recursive(graph [][]int) []int {
	N := len(graph)
	topologicalOrdering = []int{}
	edgesIn := make([]int, N)
	for i := 0; i < N; i++ {
		for _, neighbour := range graph[i] {
			edgesIn[neighbour]++
		}
	}
	noEdgesInVertices := []int{}
	for i := 0; i < N; i++ {
		if edgesIn[i] == 0 {
			noEdgesInVertices = append(noEdgesInVertices, i)
		}
	}
	colors = make([]int, N)
	for i := 0; i < N; i++ {
		colors[i] = 0 // white
	}

	for _, i := range noEdgesInVertices {
		if colors[i] == 0 {
			loopFound := DfsVisit(i, graph)
			if loopFound {
				return []int{}
			}

		}
	}
	return topologicalOrdering

}
func DfsVisit(i int, graph [][]int) bool {
	loopFound := false
	colors[i] = 1 // grey
	for _, neighbour := range graph[i] {
		if colors[neighbour] == 1 {
			// loop found, exit
			return true
		} else if colors[neighbour] == 0 {
			loopFound = DfsVisit(neighbour, graph)
		}
	}
	colors[i] = 2 //black
	topologicalOrdering = append([]int{i}, topologicalOrdering...)
	return loopFound
}
