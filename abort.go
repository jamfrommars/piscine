package piscine

func Sort(array []int) {
	for i := 0; i < 5; i++ {
		mini := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[mini] {
				mini = j
			}
		}
		array[i], array[mini] = array[mini], array[i]
	}
}

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	Sort(arr)
	return arr[2]
}
