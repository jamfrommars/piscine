package piscine

func IsPrime(nb int) bool {
	var answer int
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			answer = 1
			break
		}
	}
	if nb <= 1 {
		return false
	} else {
		if answer == 1 {
			return false
		} else {
			return true
		}
	}
}
