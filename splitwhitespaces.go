package piscine

func SplitWhiteSpaces(s string) []string {
	len := 0
	for i, ind := range s {
		if ind == ' ' && s[i+1] != ' ' {
			len++
		}
	}
	array := make([]string, len+1)
	index := 0
	result := ""
	for _, ind := range s {
		if ind == ' ' || ind == '\n' || ind == '\t' {
			if result != "" {
				array[index] = result
				result = ""
				index++
			}
		} else {
			result = result + string(ind)
		}
	}
	if result != "" {
		array[index] = result
	}
	return array
}
