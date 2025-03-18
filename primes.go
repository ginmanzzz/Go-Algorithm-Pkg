package algo

func SieveOfEratosthenes(n int) []bool {
    isPrimes := make([]bool, n+1)
    for i := 2; i <= n; i++ {
        isPrimes[i] = true
    }
    /*
        i*i <= n:
        If we want to determine whether n is prime, we only need to check whether it has [2, 3, ... sqrt(n)] divisor
        n = 36, 2*18, 3*12, 4*9, 6*6,   sqrt(36)=6, we only need to check 2, 3, 4 < 6
    */
    for i:= 2; i*i <= n; i++ {
        if isPrimes[i] {
            /*
                j = i*i
                If i = 5, 2*5, 3*5, 4*5, has been determined
            */
            for j := i*i; j <= n; j+=i {
                isPrimes[j] = false
            }
        }
    }
    return isPrimes
}
