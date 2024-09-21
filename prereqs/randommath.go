func FibonacciArray(n int) []int {
    arr := make([]int, n+1)
    arr[0] = 1
    if n == 0 {
        return arr
    }
    
    arr[1] = 1
    
    for i := 2; i <= n; i++ {
        arr[i] = arr[i-1] + arr[i-2]
    }
    
    return arr
}

func IsPerfect(n int) bool {
    return n == SumProperDivisors(n)
}

func NextPerfectNumber(n int) int {
    i := n+1
    for !IsPerfect(i) { i++ }
    return i
}

func ListMersennePrimes(n int) []int {
    var mersennePrimes []int
    for p := 1; p <= n; p++ {
        sus := Power(2, p) - 1
        if IsPrime(sus) {
            mersennePrimes = append(mersennePrimes, sus)
        }
    }
    return mersennePrimes
}

func NextTwinPrimes(n int) (int, int) {
    p := n+1
    for !(IsPrime(p) && IsPrime(p+2)) { p++ }
    return p, p+2
}
