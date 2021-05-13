package piscine

func IsPrintable(s string) bool {
	run := []rune(s)
	for i := range s {
		if run[i] >= 32 && run[i] <= 127 {
			continue
		} else {
			return false
		}
	}
	return true
}
