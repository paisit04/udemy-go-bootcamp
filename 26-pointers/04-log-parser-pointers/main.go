package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	p := newParser()
	// fmt.Printf("&main.p   : %p\n\n", &p)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {

		parsed, err := parse(&p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		update(&p, parsed)
	}

	summarize(p)

	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}

func summarize(p parser) {
	fmt.Printf("%-30s %-10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(p.domains)
	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}
	fmt.Printf("%-30s %10d\n", "TOTAL", p.total)
}
