package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for z := '0'; z <= '9'; z++ {
		for a := '0'; a <= '9'; a++ {
			y := a + 1
			for x := z; x <= '9'; x++ {
				for ; y <= '9'; y++ {
					z01.PrintRune(z)
					z01.PrintRune(a)
					z01.PrintRune(32)
					z01.PrintRune(x)
					z01.PrintRune(y)
					if z != '9' || a != '8' || x != '9' || y != '9' {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
				y = '0'
			}
		}
	}
	z01.PrintRune('\n')
}
