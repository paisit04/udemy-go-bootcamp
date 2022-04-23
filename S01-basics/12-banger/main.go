package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	msg := os.Args[1]
	// l := len(msg)
	l := utf8.RuneCountInString(msg)
	s := strings.ToUpper(msg) + strings.Repeat("!", l)
	fmt.Println(s)
}
