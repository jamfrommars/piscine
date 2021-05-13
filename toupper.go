package piscine

func ToUpper(s string) string {
	run := []rune(s)
	for i, index := range s {
		if index >= 'a' && index <= 'z' {
			run[i] = index - 32
		}
	}
	return string(run)
}
