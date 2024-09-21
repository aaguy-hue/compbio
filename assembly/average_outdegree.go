package main

// AverageOutDegree takes the adjacency list of a directed network.
// It returns the average outdegree of each node in the network.
// It does not include nodes with no outgoing edges in the average.
func AverageOutDegree(adjList map[string][]string) float64 {
	sum := 0

	for key, _ := range adjList {
		sum += len(adjList[key])
	}

	return float64(sum) / float64(len(adjList))
}
