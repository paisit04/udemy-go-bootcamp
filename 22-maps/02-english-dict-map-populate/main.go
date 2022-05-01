package main

import (
	"fmt"
	"os"
)

func main() {
	// #2A: Get the key from CLI
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}
	query := args[0]

	// #1: Empty Map Literal
	dict := map[string]string{}

	dict["up"] = "yukarı"  // adds a new pair
	dict["down"] = "aşağı" // adds a new pair
	dict["good"] = "iyi"   // #5: overwrites the value at the key: "good"
	dict["mistake"] = ""   // #8: a key with a zero-value

	for k, v := range dict {
		fmt.Printf("%q means %#v\n", k, v)
	}

	// #12: printing a map (ordered output since Go 1.12)
	fmt.Printf("%#v\n", dict)

	// #10: comma ok in a short if
	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}
	fmt.Printf("%q not found.\n", query)

	// fmt.Printf("Zero Value: %#v\n", dict)
	fmt.Printf("# of Keys : %d\n", len(dict))
}
