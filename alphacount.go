package piscine

func AlphaCount(s string) int {
	run := []rune(s)
	count := 0
	for i := range s {
		if run[i] >= 'A' && run[i] <= 'Z' || run[i] >= 'a' && run[i] <= 'z' {
			count++
		}
	}
	return count
}
