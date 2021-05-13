package piscine

func IsUpper(s string) bool {
	run := []rune(s)
	for i := range s {
		if run[i] >= 'A' && run[i] <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
