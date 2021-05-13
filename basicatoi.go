package piscine

func BasicAtoi(s string) int {
	var num int = 0
	for _, index := range s {
		counter := 0
		for i := '1'; i <= index; i++ {
			counter++
		}
		num = num*10 + counter
	}
	return num
}
