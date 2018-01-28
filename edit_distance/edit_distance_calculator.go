package edit_distance

type EditDistanceCalculator interface {
	Calculate(a, b []rune) int
	GetOperations(a, b []rune) (string, string)
}

func NewEditDistanceCalculator() EditDistanceCalculator {
	return &editDistanceCalculator{}
}

type editDistanceCalculator struct {
}

func (this editDistanceCalculator) Calculate(a, b []rune) int {
	M := len(a) + 1
	N := len(b) + 1

	previousDistance := this.initializeDistanceSlice(N)
	distance := this.initializeDistanceSlice(N)

	for i := 1; i < M; i++ {
		distance[0] = i
		for j := 1; j < N; j++ {
			distance[j] = this.Min3(previousDistance[j]+1, distance[j-1]+1, previousDistance[j-1]+this.getPenalty(a[i-1], b[j-1]))
		}
		copy(previousDistance, distance)
	}
	return distance[N-1]
}

func (this editDistanceCalculator) GetOperations(a, b []rune) (string, string) {
	M := len(a) + 1
	N := len(b) + 1

	distance := this.initializeDistanceMatrix(M, N)
	direction := this.initializeDirectionMatrix(M, N)

	for i := 1; i < M; i++ {
		for j := 1; j < N; j++ {
			distance[i][j], direction[i][j] = this.Min3WithPosition(distance[i][j-1]+1, distance[i-1][j]+1, distance[i-1][j-1]+this.getPenalty(a[i-1], b[j-1]))
		}
	}
	return this.getSequence(direction, M, N, a, b)
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
	for j := 0; j < N; j++ {
		distance[0][j] = j
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

func (this editDistanceCalculator) getSequence(direction [][]int, M, N int, a, b []rune) (string, string) {
	i := M - 1
	j := N - 1
	var letterOfFirstWord, letterOfSecondWord rune
	var firstWord, secondWord string

	for direction[i][j] != 0 && i+j != 0 {
		currentDirection := direction[i][j]
		i, j, letterOfFirstWord, letterOfSecondWord = this.performStepBack(currentDirection, i, j, a, b)
		firstWord = string(letterOfFirstWord) + firstWord
		secondWord = string(letterOfSecondWord) + secondWord
	}
	return firstWord, secondWord
}

func (editDistanceCalculator) performStepBack(currentDirection, i, j int, a, b []rune) (int, int, rune, rune) {
	switch currentDirection {
	case 1:
		return i, j - 1, '-', b[j-1]
	case 2:
		return i - 1, j, a[i-1], '-'
	default:
		return i - 1, j - 1, a[i-1], b[i-1]
	}
}

func (this editDistanceCalculator) initializeDirectionMatrix(M, N int) [][]int {
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
func (editDistanceCalculator) initializeDistanceSlice(N int) []int {
	slice := make([]int, N)
	for j := 0; j < N; j++ {
		slice[j] = j
	}
	return slice
}
