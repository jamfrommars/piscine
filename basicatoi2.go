package piscine

func BasicAtoi2(s string) int {
	var num int = 0
	for _, index := range s {
		if index >= '0' && index <= '9' {
			counter := 0
			for i := '1'; i <= index; i++ {
				counter++
			}
			num = num*10 + counter
		} else {
			return 0
		}
	}
	return num
}
