package piscine

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	num := nb
	for !isPrime(num) {
		num = num + 1
	}
	return num
}
