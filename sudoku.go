package main

import ( 
"os"
"github.com/01-edu/z01"
"fmt"
)

func main() {
	arguments := os.Args[1:]
	if ValidNumbers(arguments) == true {
	sudoku := [9][9]rune{}
	sudoku = Display(sudoku, arguments)
	OutputSudoku(sudoku)
	} else {
		fmt.Println("Error")
	}
	}

func OutputSudoku(sudoku [9][9]rune)  {
	if Solvable(&sudoku) == true {
		for x := 0; x < 9; x++ {
			for y := 0; y < 9; y++ {
				if y<8 {
					z01.PrintRune(sudoku[x][y])
					z01.PrintRune(' ')
				}else {
					z01.PrintRune(sudoku[x][y])
					}
				}
					z01.PrintRune('\n')
			}
		} else {
			fmt.Println("Error")
		}
		}

func Display(sudoku [9][9]rune, args []string) [9][9]rune {
	for i := range args {
		for j := range args[i] {
			sudoku[i][j] = rune(args[i][j])
		}
	}
	return sudoku
}

func FullBlock(sudoku *[9][9]rune, x int, y int, s rune) bool {
	horizontal := x / 3
	vertical := y / 3

	for k := 3 * horizontal; k < 3*(horizontal+1); k++ {
		for l := 3 * vertical; l < 3*(vertical+1); l++ {
			if s == sudoku[l][k] {
				return false
			}
		}
	}
return true
}
func DotCheck(sudoku *[9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == '.' {
				return true
			}
		}
	}
	return false
}
func Validity(sudoku *[9][9]rune, x int, y int, s rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
		if s == sudoku[i][x] {
			return false
		}
		if s == sudoku[y][j] {
			return false
		}
	}
}
return true
}
func ValidNumbers(args []string) bool {
	if len(args) !=9 {
            return false
	   }
	
	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9  {
			return false
		}
	}
	for i := 0; i < len(args); i++ {
		for _, index := range args[i] {
			if index == 47 || index == 48 || index < 46 || index > 57{ //в след раз заменить аски
				return false
			}
		}
	}
        return true
}

func Solvable(sudoku *[9][9]rune) bool {
	if DotCheck(sudoku)==false {
		return true
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if sudoku[y][x] == '.' {
				for k := '1'; k <= '9'; k++ {
					if Validity(sudoku, x, y, k) && FullBlock(sudoku, x , y , k)  {
						sudoku[y][x] = k
						if Solvable(sudoku) {
							return true
						}
					}
					sudoku[y][x] = '.'
				}
				return false
			}
		}
	}
	return false
}