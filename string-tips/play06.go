package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const cheers = `Cheers🍻!`
	for i, w := 0, 0; i < len(cheers); i += w {
		runeValue, width := utf8.DecodeRuneInString(cheers[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}
