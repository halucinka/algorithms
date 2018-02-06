package graph_as_matrix

func GetDistancesFromStart(graph [][]int, start int) []int {
	queue := []int{start}
	usedVertices := map[int]bool{start: true}
	N := len(graph)
	distances := make([]int, N)

	for len(queue) > 0 {
		//pop
		u := queue[0]
		queue = queue[1:]

		// for all neighbours
		for neighbour, value := range graph[u] {
			if _, ok := usedVertices[neighbour]; !ok && value != 0 {
				usedVertices[neighbour] = true
				distances[neighbour] = distances[u] + 1

				//push
				queue = append(queue, neighbour)
			}
		}
	}

	return distances
}

func GetNumberOfComponents(graph [][]int) int {
	notUsedVertices := map[int]bool{}
	for i := 0; i < len(graph); i++ {
		notUsedVertices[i] = true
	}
	components := 0
	for len(notUsedVertices) > 0 {
		start := getKeyFromMap(notUsedVertices)
		notUsedVertices = resolveComponent(start, graph, notUsedVertices)
		components++
	}
	return components

}

func resolveComponent(start int, graph [][]int, notUsedVertices map[int]bool) map[int]bool {
	queue := []int{start}
	delete(notUsedVertices, start)
	for len(queue) > 0 {
		//pop
		u := queue[0]
		queue = queue[1:]

		// for all neighbours
		for neighbour, value := range graph[u] {
			if _, ok := notUsedVertices[neighbour]; ok && value != 0 {
				delete(notUsedVertices, neighbour)

				//push
				queue = append(queue, neighbour)
			}
		}
	}
	return notUsedVertices
}

func getKeyFromMap(m map[int]bool) int {
	for k := range m {
		return k
	}
	return -1
}
