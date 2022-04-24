package main

import (
	"fmt"
	"time"
)

func main() {
	switch h := time.Now().Hour(); {
	case h >= 18:
		fmt.Println("Good Evening")
	case h >= 12:
		fmt.Println("Good Afternoon")
	case h >= 6:
		fmt.Println("Good Morning")
	default:
		fmt.Println("Good Night")
	}
}
