package student

import "fmt"

//Raid1a is drawing square area
func Raid1a(x, y int) {

	for col := 0; col < x; col++ {
		for row := 0; row < y; row++ {
			if col == 0 && row == 0 || col == x-1 && row == y-1 || col == 0 && row == y-1 || col == x-1 && row == 0 {
				fmt.Print("o")
			} else if col == 0 || col == x-1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		if col == x-1 {
			fmt.Print("\n")
		}
	}
}
