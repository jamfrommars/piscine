package piscine

func NRune(s string, n int) rune {
	run := []rune(s)
	count := 0
	for i := range run {
		count = i + 1
	}
	if n > 0 && n <= count {
		return run[n-1]
	} else {
		return 0
	}
}
