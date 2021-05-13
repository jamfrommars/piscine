package piscine

func BasicJoin(elems []string) string {
	str := ""
	for _, index := range elems {
		str = str + index
	}
	return str
}
