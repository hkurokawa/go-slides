package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	const cheers = "\t\tCheers\t\U0001F37B!\t\n"
	fmt.Println("Before TrimFunc:")
	fmt.Println(cheers)
	s := strings.TrimFunc(cheers, unicode.IsControl)
	fmt.Println("After TrimFunc:")
	fmt.Println(s)
}
