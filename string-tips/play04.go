package main

import "fmt"
import "unicode/utf8"

func main() {
	const sample = `Hello, 🍺`

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("len():")
	fmt.Println(len(sample))

	fmt.Println("utf8.RuneCountInString():")
	fmt.Println(utf8.RuneCountInString(sample))

	fmt.Println("range loop:")
	for _, c := range sample {
		fmt.Printf("%+q ", c)
	}
	fmt.Println()
}
