package main

import "fmt"

func main() {
	s1 := "Hello, 世界"
	fmt.Println(len(s1)) // -> 13
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%x ", s1[i])
	}
	fmt.Println()

	s2 := s1[7:]
	fmt.Println(s2) // -> "世界"

	// Below code causes a compile error!
	// s1[0] = 'h'
}
