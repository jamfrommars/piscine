package piscine

func ToLower(s string) string {
	run := []rune(s)
	for i, index := range s {
		if index >= 'A' && index <= 'Z' {
			run[i] = index + 32
		}
	}
	return string(run)
}
