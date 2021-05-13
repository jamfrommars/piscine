package piscine

func TrimAtoi(s string) int {
	var result int
	var neglect int
	for _, index := range s {
		count := 0
		if result == 0 && index == '-' {
			neglect = 1
		}
		if index >= '0' && index <= '9' {
			for i := '0'; i < index; i++ {
				count++
			}
			result = 10*result + count
		}
	}
	if neglect == 1 {
		return -result
	}
	return result
}
