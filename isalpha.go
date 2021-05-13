package piscine

func IsAlpha(s string) bool {
	run := []rune(s)
	for i := range s {
		if run[i] >= 'A' && run[i] <= 'Z' || run[i] >= 'a' && run[i] <= 'z' || run[i] >= '0' && run[i] <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
