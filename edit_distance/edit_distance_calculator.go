package edit_distance


type EditDistanceCalculator interface {
    Calculate(a, b []rune) int
}

func NewEditDistanceCalculator() EditDistanceCalculator {
    return &editDistanceCalculator{}
}

type editDistanceCalculator struct {
}

func (this editDistanceCalculator) Calculate(a, b []rune) int {
    M := len(a)+1
    N := len(b)+1

    distance := this.declare2DSlice(M, N)

    for i := 0; i < M; i++ {
        distance[i][0] = i
    }
    for i := 0; i < N; i++ {
        distance[0][i] = i
    }

    for i := 0; i < M-1; i ++ {
        for j := 0; j < N-1; j++ {
            distance[i+1][j+1] = this.Min3(distance[i][j+1]+1, distance[i+1][j]+1, distance[i][j]+this.getPenalty(a[i], b[j]))
        }
    }
    return distance[M-1][N-1]
}

func (this editDistanceCalculator) Min3(x, y, z int) int {
    return this.min(this.min(x, y), z)
}

func (editDistanceCalculator) declare2DSlice(M int, N int) [][]int {
    matrix := make([][]int, M)
    for i := 0; i < M; i++ {
        matrix[i] = make([]int, N)
    }
    return matrix
}

func (editDistanceCalculator) getPenalty( a, b rune) int{
    if a == b {
        return 0
    }
    return 1
}

func (editDistanceCalculator) min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
