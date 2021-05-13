package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	arr := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		arr[i] = i + min
	}
	return arr
}
