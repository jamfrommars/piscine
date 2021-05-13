package piscine

func IterativeFactorial(nb int) int {
	if nb >= 0 && nb < 25 {
		result := 1
		for i := 1; i < nb+1; i++ {
			result = result * i
		}
		return result
	} else {
		return 0
	}
}
