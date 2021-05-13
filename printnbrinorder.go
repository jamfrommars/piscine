package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var arr []int
	var minValue int
	counter := 0
	if n == 0 {
		z01.PrintRune('0')
	}
	for n != 0 {
		remain := n % 10
		n /= 10
		arr = append(arr, remain)
	}
	for count := range arr {
		counter = count + 1
	}
	for i := 0; i < counter-1; i++ {
		for j := 0; j < counter-i-1; j++ {
			if arr[j] > arr[j+1] {
				minValue = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = minValue
			}
		}
	}
	for i := 0; i < counter; i++ {
		z01.PrintRune(rune(arr[i] + 48))
	}
}
