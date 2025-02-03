package main

type Edge struct {
	From   Coordinate // Coord 1 in (row, col) format
	To     Coordinate // Coord 2 in (row, col) format
	Weight int        // The weight of the edge
}

type Coordinate struct {
	Row int
	Col int
}

// this just stores edges as a map of maps instead of a list of edges for O(1) lookups
type EdgeMap map[Coordinate]map[Coordinate]int

// ManhattanTourist takes in the rows and columns in the matrix and the edge weights.
// It returns a path through the matrix that maximizes the sum of the edge weights.
func ManhattanTourist(n, m int, edges []Edge) []Coordinate {
	edgeMap := BuildEdgeMap(edges)

	// find the scored matrix
	scored := ManhattanTouristScore(n, m, edgeMap)

	// backtracking through the matrix to find the path
	// basically we're retracing our steps, looking at which direction we came from by looking at
	// which edge we took to get to the current node
	path := []Coordinate{}
	path = append(path, Coordinate{n - 1, m - 1})
	i, j := n-1, m-1
	for i > 0 || j > 0 { // while not at the top left corner
		if i != 0 && j == 0 { // if we're at the left edge, go up
			path = append(path, Coordinate{i - 1, j})
			i--
		} else if i == 0 && j != 0 { // if we're at the top edge, go left
			path = append(path, Coordinate{i, j - 1})
			j--
		} else {
			up := scored[i-1][j] + FindEdge(edgeMap, Coordinate{i - 1, j}, Coordinate{i, j})
			left := scored[i][j-1] + FindEdge(edgeMap, Coordinate{i, j - 1}, Coordinate{i, j})

			if up == scored[i][j] {
				path = append(path, Coordinate{i - 1, j})
				i--
			} else if left == scored[i][j] {
				path = append(path, Coordinate{i, j - 1})
				j--
			} else {
				// I'm gonna crash out if this case somehow happens
			}
		}
	}

	return path
}

// ManhattanTouristScore takes in the rows and columns in the matrix and the edge weights.
// It returns the scored matrix.
func ManhattanTouristScore(n, m int, edges EdgeMap) [][]int {
	// init matrix with 0s
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}

	// assign the top row scores by looping through the cols
	for i := 1; i < m; i++ {
		arr[0][i] = arr[0][i-1] + FindEdge(edges, Coordinate{0, i - 1}, Coordinate{0, i})
	}

	// assign the left col scores by looping through the rows
	for i := 1; i < n; i++ {
		arr[i][0] = arr[i-1][0] + FindEdge(edges, Coordinate{i - 1, 0}, Coordinate{i, 0})
	}

	// assign score for each node
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			up := arr[i-1][j] + FindEdge(edges, Coordinate{i - 1, j}, Coordinate{i, j})
			left := arr[i][j-1] + FindEdge(edges, Coordinate{i, j - 1}, Coordinate{i, j})

			arr[i][j] = Max(up, left)
		}
	}

	return arr
}

// Stores the edges as a map of edges for faster lookups (we won't need to do a linear search each time)
func BuildEdgeMap(edges []Edge) EdgeMap {
	edgeMap := make(EdgeMap)
	for _, edge := range edges {
		if edgeMap[edge.From] == nil {
			edgeMap[edge.From] = make(map[Coordinate]int)
		}
		edgeMap[edge.From][edge.To] = edge.Weight
	}
	return edgeMap
}

// Determines whether an edge exists in a slice of edges
func FindEdge(edgeMap EdgeMap, from, to Coordinate) int {
	if weight, exists := edgeMap[from][to]; exists {
		return weight
	}
	return 0
}
