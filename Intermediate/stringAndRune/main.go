package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "Hello"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d ", s[i])
	}

	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")

	for i, w := 0, 0; i < len(s); i += w {

		// This line uses the utf8.DecodeRuneInString function to decode the next rune in the string s starting from index i. It returns two values: runeValue is the decoded rune, and width is the number of bytes used by the rune. The loop processes runes one by one in each iteration.

		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		fmt.Println("w : ", w)

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 'e' {
		fmt.Println("found ee")
	} else {
		fmt.Println("found so sua")
	}
}
