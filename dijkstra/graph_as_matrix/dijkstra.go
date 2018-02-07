package graph_as_matrix

func GetDistancesFromStart(graph [][]int, start int) []int {
	infinity := 100000000
	N := len(graph)

	notUsedVertices := map[int]bool{}
	for index := 0; index < N; index++ {
		notUsedVertices[index] = true
	}
	distances := make([]int, N)
	for index := 0; index < N; index++ {
		distances[index] = infinity
	}
	distances[start] = 0
	for len(notUsedVertices) > 0 {
		// extractMin
		minimum := infinity
		u := -1
		for v := range notUsedVertices {
			if distances[v] < minimum {
				minimum = distances[v]
				u = v
			}
		}
		if u == -1 {
			// disconnected graph
			break
		}
		delete(notUsedVertices, u)
		// for all neighbours
		for neighbour, value := range graph[u] {
			if value != 0 {
				if distances[neighbour] > distances[u]+value {
					distances[neighbour] = distances[u] + value
				}
			}
		}
	}
	return distances
}
