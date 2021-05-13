package piscine

func LastRune(s string) rune {
	run := []rune(s)
	return run[len(s)-1]
}
