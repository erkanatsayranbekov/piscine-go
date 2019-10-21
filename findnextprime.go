package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}

	}
	return true
}

func FindNextPrime(nb int) int {
	nextprime := nb - 1
	for i := nb + 1; i > nb; i++ {
		nextprime++
		if IsPrime(nextprime) {
			return nextprime
		}
	}
	return nextprime
}
