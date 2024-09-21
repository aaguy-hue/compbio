package main

func AssertSameLength(arr []string) {
	if len(arr) == 0 {
		return
	}

	for _, item := range arr {
		if len(item) != len(arr[0]) {
			panic("Array does not have items all of the same length.")
		}
	}
}
