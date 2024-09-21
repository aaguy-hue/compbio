package main

import "fmt"

func GenerateLinearCongruenceSequence(seed, a, c, m int) []int {
	var retval []int = []int{seed}
	for !HasRepeat(retval) {
		seed = (a*seed + c) % m
		retval = append(retval, seed)
	}
	return retval
}

func main() {
	fmt.Println(ComputePeriodLength(GenerateLinearCongruenceSequence(1, 5, 1, 8192)))
}
