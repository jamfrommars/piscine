package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	for i := 0; i < len(s)-len(toFind)+1; i++ {
		slack := s[i:]
		res := 0
		for i := range toFind {
			if slack[i] != toFind[i] {
				break
			}
			res++
		}
		if res == len(toFind) {
			return i
		}
	}
	return -1
}
