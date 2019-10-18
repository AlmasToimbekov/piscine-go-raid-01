package student

import (
	"fmt"

	"github.com/01-edu/z01"
)

//Raid1a is drawing square area
func Raid1a(x, y int) {

	for row := 0; row < y; row++ {
		printRow(row, x, y)
	}
}

func printRow(row int, x int, y int) {
	for colomn := 0; colomn < x; colomn++ {
		if colomn == 0 && row == 0 || colomn == x-1 && row == y-1 || colomn == 0 && row == y-1 || colomn == x-1 && row == 0 {
			fmt.Print("o")
		} else if colomn == 0 || colomn == x-1 {
			fmt.Print("|")
		} else if row == 0 || row == y-1 {
			fmt.Print("-")
		} else {
			z01.PrintRune(' ')
		}
		if colomn == x-1 {
			fmt.Print("\n")
		}
	}
}
