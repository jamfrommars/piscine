package piscine

func alpha(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	} else {
		return false
	}
}

func Capitalize(s string) string {
	arr := []rune(s)
	letter := true
	for i := 0; i < len(s); i++ {
		if alpha(arr[i]) == true && letter {
			if arr[i] >= 'a' && arr[i] <= 'z' {
				arr[i] = 'A' - 'a' + arr[i]
			}
			letter = false
		} else if arr[i] >= 'A' && arr[i] <= 'Z' {
			arr[i] = 'a' - 'A' + arr[i]
		} else if alpha(arr[i]) == false {
			letter = true
		}
	}
	return string(arr)
}
