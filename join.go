package piscine

func Join(strs []string, sep string) string {
	str := ""
	counter := 0
	for i := range strs {
		counter = i + 1
	}
	for i := range strs {
		if i == counter-1 {
			str += strs[i]
		} else {
			str += strs[i] + sep
		}
	}
	return str
}
