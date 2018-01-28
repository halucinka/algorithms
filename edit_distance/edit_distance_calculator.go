package edit_distance

type EditDistanceCalculator interface {
	Calculate(a, b []rune) int
	GetShortestSequenceOfOperations(a, b []rune) []string
}

func NewEditDistanceCalculator() EditDistanceCalculator {
	return &editDistanceCalculator{}
}

type editDistanceCalculator struct {
}

func (this editDistanceCalculator) Calculate(a, b []rune) int {
	M := len(a) + 1
	N := len(b) + 1

	distance := this.initializeDistanceMatrix(M, N)
	for i := 0; i < M-1; i++ {
		for j := 0; j < N-1; j++ {
			distance[i+1][j+1] = this.Min3(distance[i][j+1]+1, distance[i+1][j]+1, distance[i][j]+this.getPenalty(a[i], b[j]))
		}
	}
	return distance[M-1][N-1]
}

func (this editDistanceCalculator) GetShortestSequenceOfOperations(a, b []rune) []string {
	M := len(a) + 1
	N := len(b) + 1

	distance := this.initializeDistanceMatrix(M, N)
	direction := this.initializeDirection(M, N)

	for i := 0; i < M-1; i++ {
		for j := 0; j < N-1; j++ {
			distance[i+1][j+1], direction[i+1][j+1] = this.getOptimalStep(a, b, i, j, distance)
		}
	}
	return this.getSequence(direction, M, N)
}

func (this editDistanceCalculator) Min3(x, y, z int) int {
	return this.min(this.min(x, y), z)
}

func (this editDistanceCalculator) Min3WithPosition(x, y, z int) (int, int) {
	if this.min(x, y) == x {
		if this.min(x, z) == x {
			return x, 1
		}
	}
	if this.min(y, z) == y {
		return y, 2
	}
	return z, 3
}

func (this editDistanceCalculator) getOptimalStep(a, b []rune, i, j int, distance [][]int) (int, int) {
	penalty := this.getPenalty(a[i], b[j])
	minimumSteps, previousStep := this.Min3WithPosition(distance[i+1][j]+1, distance[i][j+1]+1, distance[i][j]+penalty)
	if this.charactersAreIdentical(previousStep, penalty) {
		previousStep = 4
	}
	return minimumSteps, previousStep
}

func (this editDistanceCalculator) charactersAreIdentical(position int, penalty int) bool {
	return position == 3 && penalty == 0
}

func (editDistanceCalculator) min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (this editDistanceCalculator) initializeDistanceMatrix(M int, N int) [][]int {
	distance := this.declare2DSlice(M, N)
	for i := 0; i < M; i++ {
		distance[i][0] = i
	}
	for i := 0; i < N; i++ {
		distance[0][i] = i
	}
	return distance
}

func (editDistanceCalculator) declare2DSlice(M int, N int) [][]int {
	matrix := make([][]int, M)
	for i := 0; i < M; i++ {
		matrix[i] = make([]int, N)
	}
	return matrix
}

func (editDistanceCalculator) getPenalty(a, b rune) int {
	if a == b {
		return 0
	}
	return 1
}

func (this editDistanceCalculator) getSequence(direction [][]int, M, N int) []string {
	i := M - 1
	j := N - 1
	sequence := []string{}
	for direction[i][j] != 0 {
		currentDirection := direction[i][j]
		a, b, operation := this.performStepBack(currentDirection, i, j)
		if operation != "" {
			sequence = append([]string{operation}, sequence...)
		}
		if a+b == 0 {
			break
		}
		i = a
		j = b
	}
	return sequence
}

func (editDistanceCalculator) performStepBack(currentDirection, i, j int) (int, int, string) {
	switch currentDirection {
	case 1:
		return i, j - 1, "add"
	case 2:
		return i - 1, j, "delete"
	case 3:
		return i - 1, j - 1, "mutate"
	case 4:
		return i - 1, j - 1, ""
	default:
		return 0, 0, ""
	}
}
func (this editDistanceCalculator) initializeDirection(M, N int) [][]int {
	direction := this.declare2DSlice(M, N)
	direction[0][0] = 0
	for i := 1; i < M; i++ {
		direction[i][0] = 2
	}
	for i := 1; i < N; i++ {
		direction[0][i] = 1
	}
	return direction
}
