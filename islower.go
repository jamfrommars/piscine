package piscine

func IsLower(s string) bool {
	run := []rune(s)
	for i := range s {
		if run[i] >= 'a' && run[i] <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
