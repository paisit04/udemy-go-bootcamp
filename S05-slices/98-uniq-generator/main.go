package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	max, _ := strconv.Atoi(os.Args[1])

	var uniques []int

loop:
	for len(uniques) < max {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		uniques = append(uniques, n)
	}

	fmt.Println("\n\nuniques:", uniques)

	sort.Ints(uniques)
	fmt.Println("\n\nsorted:", uniques)

	nums := [5]int{5, 4, 3, 2, 1}
	sort.Ints(nums[:])
	fmt.Println("\n\nnums:", nums)
}
