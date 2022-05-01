package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("gimme somethin' to mask!")
		return
	}

	const (
		link  = "http://"
		nlink = len(link)
		mask  = '*'
	)

	var (
		text = args[0]
		size = len(text)
		buf  = make([]byte, 0, size)

		in bool
	)

	for i := 0; i < size; i++ {
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			in = true

			buf = append(buf, link...)
			i += nlink
		}

		c := text[i]

		switch c {
		case ' ', '\t', '\n':
			in = false
		}

		if in {
			c = mask
		}
		buf = append(buf, c)
	}

	// print out the buffer as text (string)
	fmt.Println(string(buf))
}
