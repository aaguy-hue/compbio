package main

func IsPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func LargestPrimeFactor(n int) int {
	for i := n; i > 1; i-- {
		if n%i == 0 && IsPrime(i) {
			return i
		}
	}

	return 1 // should never happen since n >= 2
}

// func main() {
// 	var inp int
// 	fmt.Print("Enter a number: ")
// 	fmt.Scan(&inp)
// 	fmt.Println("The largest prime factor is", LargestPrimeFactor(inp))
// }
