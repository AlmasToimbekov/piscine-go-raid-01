package student

import (
	"github.com/01-edu/z01"
)

//Raid1c is drawing square area from letters
func Raid1c(x, y int) {

	for row := 0; row < y; row++ {
		printRowForC(row, x, y)
	}
}

func printRowForC(row int, x int, y int) {
	for colomn := 0; colomn < x; colomn++ {
		if row == 0 && (colomn == x-1 || colomn == 0) {
			z01.PrintRune('A')
		} else if row == y-1 && (colomn == x-1 || colomn == 0) {
			z01.PrintRune('C')
		} else if colomn == 0 || colomn == x-1 || row == 0 || row == y-1 {
			z01.PrintRune('B')
		} else {
			z01.PrintRune(' ')
		}
		if colomn == x-1 {
			z01.PrintRune(10)
		}
	}
}
