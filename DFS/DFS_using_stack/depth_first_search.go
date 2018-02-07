package DFS_using_stack

type Vertex struct {
	id                       int
	firstNotVisitedNeighbour int
}

func GetNumberOfComponents(graph [][]int) int {
	componentsCount := DFS(graph)
	return componentsCount
}

func DFS(graph [][]int) int {
	N := len(graph)
	colors := make([]int, N)
	for i := 0; i < N; i++ {
		colors[i] = 0 // white
	}
	componentsCount := 0

	for i := 0; i < N; i++ {
		if colors[i] == 0 {
			componentsCount++
			stack := []*Vertex{{id: i, firstNotVisitedNeighbour: 0}}
			colors[i] = 1 //grey

			for len(stack) > 0 {
				//loop up
				u := stack[len(stack)-1]

				if u.firstNotVisitedNeighbour >= len(graph[u.id]) {
					stack = stack[0 : len(stack)-1]
					colors[u.id] = 2 // black
				} else {
					neighbour := graph[u.id][u.firstNotVisitedNeighbour]
					u.firstNotVisitedNeighbour++
					if colors[neighbour] == 0 {
						colors[neighbour] = 1
						stack = append(stack, &Vertex{id: neighbour, firstNotVisitedNeighbour: 0})
					}
				}
			}
		}
	}
	return componentsCount
}
