package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, index := range a {
		for _, i := range index {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
