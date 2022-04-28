package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	// create the digits

	// clear the screen
	screen.Clear()

	// draw loop
	for {
		// move the cursor to top-left
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		ssec := now.Nanosecond() / 1e8

		// fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

		// update the clock array by copying the digits into it
		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
			dot,
			digits[ssec],
		}

		if sec%10 == 0 {
			clock = alarm
		}

		// draw the Clock using the clock array
		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if (digit == colon || digit == dot) && sec%2 == 0 {
					next = "   "
				}

				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		fmt.Println()

		const splitSecond = time.Second / 10
		time.Sleep(splitSecond)
	}

}
