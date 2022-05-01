package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}
	query := args[0]

	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mükemmel",
		"awesome": "mükemmel", // #5
	}

	delete(dict, "awesome") // #6
	delete(dict, "awesome") // #7: no-op
	delete(dict, "notexisting")
	// dict = nil // #8
	// for k := range dict { // #9
	// 	delete(dict, k)
	// }

	// turkish := dict
	turkish := make(map[string]string, len(dict))
	// fmt.Println("len(turkish map)", len(turkish))

	for k, v := range dict {
		turkish[v] = k
	}

	// fmt.Printf("english: %q\nturkish: %q\n", dict, turkish)

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	if value, ok := turkish[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	fmt.Printf("%q not found.\n", query)
}
