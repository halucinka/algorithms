package DFS_using_recursion

var stepNumber int
var discoveryOrderingIn = []int{}
var discoveryOrderingOut = []int{}
var parent = []int{}
var colors = []int{}

func GetNumberOfComponents_recursive(graph [][]int) int {
	componentsCount, _, _, _ := DFS_recursive(graph)
	return componentsCount
}

func DFS_recursive(graph [][]int) (int, []int, []int, []int) {
	N := len(graph)
	colors = make([]int, N)
	discoveryOrderingIn = make([]int, N)
	discoveryOrderingOut = make([]int, N)
	parent = make([]int, N)
	for i := 0; i < N; i++ {
		colors[i] = 0 // white
		discoveryOrderingIn[i] = -1
		discoveryOrderingOut[i] = -1
		parent[i] = -1
	}
	stepNumber = 0
	componentsCount := 0

	for i := 0; i < N; i++ {
		if colors[i] == 0 {
			componentsCount++
			DfsVisit(i, graph)
		}
	}
	return componentsCount, discoveryOrderingIn, discoveryOrderingOut, parent

}
func DfsVisit(i int, graph [][]int) {
	stepNumber++
	discoveryOrderingIn[i] = stepNumber
	colors[i] = 1 // grey
	for _, neighbour := range graph[i] {
		if colors[neighbour] == 0 {
			parent[neighbour] = i
			DfsVisit(neighbour, graph)
		}
	}
	colors[i] = 2 //black
	stepNumber++
	discoveryOrderingOut[i] = stepNumber
}
