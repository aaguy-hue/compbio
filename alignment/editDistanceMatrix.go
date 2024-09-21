package main

// EditDistanceMatrix takes as input a slice of strings patterns.
// It returns a matrix whose (i,j)th entry is the edit distance between
// the i-th and j-th strings in patterns.
func EditDistanceMatrix(patterns []string) [][]int {
	var matrix [][]int = make([][]int, len(patterns))
	for i := 0; i < len(patterns); i++ {
		matrix[i] = make([]int, len(patterns))
	}

	for r := 0; r < len(patterns); r++ {
		for c := 0; c < r; c++ {
			dist := EditDistance(patterns[r], patterns[c])
			matrix[r][c] = dist
			matrix[c][r] = dist
		}
	}

	return matrix
}
