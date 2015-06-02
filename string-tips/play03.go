package main

import "fmt"

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Byte loop printing with %q:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%q ", sample[i])
	}
	fmt.Println()
}
