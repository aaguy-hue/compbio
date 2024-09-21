package main

func RichnessMap(allMaps map[string](map[string]int)) map[string]int {
	richnessMap := make(map[string]int)
	for k, v := range allMaps {
		richnessMap[k] = Richness(v)
	}
	return richnessMap
}

func Richness(sample map[string]int) int {
	// richness = how many distinct species, sample = freq table
	var richness int
	for _, v := range sample {
		if v > 0 {
			richness += 1
		}
	}
	return richness
}

func SumOfValues(dict map[string]int) int {
	// adds the values in the table
	var sum int = 0
	for _, val := range dict {
		sum += val
	}
	return sum
}
