package main

import "fmt"

func main() {
	// args := os.Args[1:]
	// if len(args) != 1 {
	// 	fmt.Println("[english word] -> [turkish word]")
	// 	return
	// }
	// query := args[0]

	var dict map[string]string

	// #2: Map retrieval is O(1) — on average.
	key := "good"

	value := dict[key]
	fmt.Printf("%q means %#v\n", key, value)

	// #1B
	fmt.Printf("# of Keys: %d\n", len(dict))
}
