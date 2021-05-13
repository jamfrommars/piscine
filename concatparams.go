package piscine

func ConcatParams(args []string) string {
	result := ""
	for i, index := range args {
		result += string(index)
		if i != len(args)-1 {
			result += "\n"
		}
	}
	return result
}
