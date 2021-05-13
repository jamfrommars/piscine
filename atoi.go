package piscine

func Atoi(s string) int {
	var num int = 0
	var sign int = 0
	for a, index := range s {
		if a == 0 {
			if index == '-' || index == '+' {
				if s[0] == '-' {
					sign++
				}
				continue
			}
		}
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
	if sign == 0 {
		return num
	} else {
		return -num
	}
}
