package main

import "sort"

func BetaDiversityMatrix(allMaps map[string](map[string]int), distanceMetric string) ([]string, [][]float64) {
	distances := make([][]float64, len(allMaps))

	for i := 0; i < len(allMaps); i++ {
		distances[i] = make([]float64, len(allMaps))
	}

	names := make([]string, len(allMaps))

	var i int = 0
	for key, _ := range allMaps {
		names[i] = key
		i++
	}

	sort.Strings(names)

	for i := 0; i < len(names); i++ {
		sample1 := allMaps[names[i]]
		for j := i + 1; j < len(names); j++ {
			// since values in the matrix over the diangonal are symmetrical, we don't need to calculate them if i < j
			// we skip the values on the diagonal since the distance will be 0 and the default value for floats is 0
			sample2 := allMaps[names[j]]
			distances[i][j] = GetDistance(sample1, sample2, distanceMetric)
			distances[j][i] = distances[i][j]
		}
	}

	return names, distances
}

func GetDistance(map1, map2 map[string]int, distanceMetric string) float64 {
	if distanceMetric == "Bray-Curtis" {
		return BrayCurtisDistance(map1, map2)
	} else if distanceMetric == "Jaccard" {
		return JaccardDistance(map1, map2)
	} else {
		panic("Invalid distance metric!")
	}
}

// JaccardDistance takes two frequency maps and returns the Jaccard
// distance between them.
func JaccardDistance(map1 map[string]int, map2 map[string]int) float64 {
	c := SumOfMinima(map1, map2)
	t := SumOfMaxima(map1, map2)
	j := 1 - (float64(c) / float64(t))
	return j
}

// BrayCurtisDistance takes two frequency maps and returns the Bray-Curtis
// distance between them.
func BrayCurtisDistance(map1 map[string]int, map2 map[string]int) float64 {
	c := SumOfMinima(map1, map2)
	s1 := SampleTotal(map1)
	s2 := SampleTotal(map2)

	//don't allow the situation in which we have zero richness.
	if s1 == 0 || s2 == 0 {
		panic("Error: sample given to BrayCurtisDistance() has no positive values.")
	}

	av := Average(float64(s1), float64(s2))
	return 1 - (float64(c) / av)
}

// SumOfMinima takes two frequency maps as input.
// It identifies the keys that are shared by the two frequency maps
// and returns the sum of the corresponding values.
func SumOfMinima(map1 map[string]int, map2 map[string]int) int {
	c := 0

	for key := range map1 {
		_, exists := map2[key]
		if exists {
			c += Min2(map1[key], map2[key])
		}
	}

	return c
}

// Min2 takes two integers and returns their minimum.
func Min2(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func SumOfMaxima(map1 map[string]int, map2 map[string]int) int {
	c := 0

	for key := range map2 {
		_, exists := map1[key]
		if exists {
			c += Max2(map1[key], map2[key])
		} else {
			c += map2[key]
		}
	}
	for key := range map1 {
		_, exists := map2[key]
		if !exists {
			c += map1[key]
		}
	}

	// panic if c is equal to zero since we don't want 0/0
	if c == 0 {
		panic("Error: no species common to two maps given to SumOfMaxima")
	}

	return c
}

func Max2(n1, n2 int) int {
	if n1 < n2 {
		return n2
	}
	return n1
}

// Average takes two floats and returns their average.
func Average(x, y float64) float64 {
	return (x + y) / 2.0
}
