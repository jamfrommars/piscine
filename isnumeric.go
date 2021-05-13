package piscine

func IsNumeric(s string) bool {
	run := []rune(s)
	for i := range s {
		if run[i] >= '0' && run[i] <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
