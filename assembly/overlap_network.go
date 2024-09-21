package main

// MakeOverlapNetwork() takes a slice of reads with match, mismatch, gap and a threshold.
// It returns adjacency lists of reads; edges are only included
// in the overlap graph is the alignment score of the two reads is at least the threshold.
func MakeOverlapNetwork(reads []string, match, mismatch, gap, threshold float64) map[string][]string {
	//Initialize an adjacency list to represent the overlap graph.
	adjacencyList := make(map[string][]string)
	mtx := OverlapScoringMatrix(reads, match, mismatch, gap)
	binarizedMtx := BinarizeMatrix(mtx, threshold)
	rows := len(binarizedMtx)
	cols := len(binarizedMtx[0])

	for i := 0; i < rows; i++ {
		node1 := reads[i]
		adjacencyList[node1] = make([]string, 0) // ensure that a key will be there, even if there are no matches
		for j := 0; j < cols; j++ {
			if binarizedMtx[i][j] == 1 {
				node2 := reads[j]
				adjacencyList[node1] = append(adjacencyList[node1], node2)
			}
		}
	}

	return adjacencyList
}
