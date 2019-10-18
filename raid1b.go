package student

import (
	"github.com/01-edu/z01"
)

//Raid1b is drawing rhombus area
func Raid1b(x, y int) {

	for row := 0; row < y; row++ {
		printRowb(row, x, y)
	}
}

func printRowb(row int, x int, y int) {
	for colomn := 0; colomn < x; colomn++ {
		if colomn == 0 && row == 0 || colomn == x-1 && row == x-y {
			z01.PrintRune('/')
		} else if colomn == x-1 && row == 0 || colomn == 0 && row == y-1 {
			z01.PrintRune('\\')
		} else if colomn == 0 || colomn == x-1 || row == 0 || row == y-1 {
			z01.PrintRune('*')
		} else {
			z01.PrintRune(' ')
		}
		if colomn == x-1 {
			z01.PrintRune(10)
		}
	}
}
