package piscine

func AppendRange(min, max int) []int {
	var arr []int
	for i := min - 1; i < max-1; i++ {
		arr = append(arr, i+1)
	}
	return arr
}
