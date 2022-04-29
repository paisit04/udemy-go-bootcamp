// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Correct the lyric
//
//  You have a slice that contains the words of Beatles'
//  legendary song: Yesterday. However, the order of the
//  words are incorrect.
//
// CURRENT OUTPUT
//
//  [all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
// EXPECTED OUTPUT
//
//  [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday]
//
//
// STEPS
//
//  INITIAL SLICE:
//    [all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
//
//  1. Prepend "yesterday" to the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
//
//  2. Put the words to the correct positions in the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday]
//
//
//  3. Print the `lyric` slice.
//
//
// BONUS
//
//   + Think about when does the append allocate a new backing array.
//
//   + Check whether your conclusions are correct.
//
//
// HINTS
//
//   If you get stuck, check out the hints.md file.
//
// ---------------------------------------------------------

func main() {
	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)

	// ADD YOUR CODE BELOW:
	lyric = append([]string{"yesterday"}, lyric...)

	// ------------------------------------------------------------------------
	// #2: Put the words to the correct positions in the `lyric` slice.
	// ------------------------------------------------------------------------

	//
	// yesterday all my troubles seemed so far away oh I believe in yesterday
	//                                              |               |
	//                                              v               v
	//                                       index: 8          pos: 13
	//
	const N, M = 8, 13
	// --- BEFORE ---
	//
	// yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay
	//
	// --- AFTER ---
	// yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay oh I believe in yesterday
	//                                              ^
	//
	//                                              |
	//                                              +--- duplicate
	//
	// (warning: allocates a new backing array)
	lyric = append(lyric, lyric[N:M]...)
	//
	// --- BEFORE ---
	// yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay oh I believe in yesterday
	//
	// --- AFTER ---
	// yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday
	//
	// (does not allocate a new backing array because cap(lyric[:N]) > M)
	//
	lyric = append(lyric[:N], lyric[M:]...)
	fmt.Printf("%s\n", lyric)
}
