package student

import "github.com/01-edu/z01"

//Raid1e prints a valid square of width x and of height y
func Raid1e(x, y int) {
	for row := 0; row < y; row++ {
		printRowForE(row, x, y)
	}
}

func printRowForE(row int, x int, y int) {
	for colomn := 0; colomn < x; colomn++ {
		if colomn == 0 && row == 0 || colomn == x-1 && row == y-1 && x != 1 && y > 1 {
			z01.PrintRune('A')
		} else if colomn == x-1 && row == 0 || colomn == 0 && row == y-1 {
			z01.PrintRune('C')
		} else if colomn == 0 || row == 0 || colomn == x-1 || row == y-1 {
			z01.PrintRune('B')
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(10)
}
