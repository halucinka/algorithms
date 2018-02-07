package DFS_using_stack

func GetNumberOfComponents(graph [][]int) int {
	componentsCount := DFS(graph)
	return componentsCount
}

func DFS(graph [][]int) int {
	N := len(graph)
	visited := make([]bool, N)
	for i := 0; i < N; i++ {
		visited[i] = false
	}
	stepNumber := 0
	componentsCount := 0

	for i := 0; i < N; i++ {
		if !visited[i] {
			componentsCount++
			stack := []int{i}
			visited[i] = true

			for len(stack) > 0 {
				//loop up
				u := stack[len(stack)-1]

				stepNumber++
				someNeighbourIsNotVisited := false
				for _, neighbour := range graph[u] {
					if !visited[neighbour] {
						// we pick one neighbour at a time, so every node will be in stack at most once
						visited[neighbour] = true
						stack = append(stack, neighbour)
						someNeighbourIsNotVisited = true
						break
					}
				}
				if !someNeighbourIsNotVisited {
					// leaving node = node becomes black
					stack = stack[0 : len(stack)-1]
				}
			}
		}
	}
	return componentsCount
}
