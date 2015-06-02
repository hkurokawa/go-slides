package main

import (
	"bytes"
	"fmt"
)

var list = [...]string{
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
}

func main() {
	fmt.Println(concat())
	fmt.Println(appendBytes())
	fmt.Println(useBytesBuff())
}

func concat() string {
	s := ""
	for _, v := range list {
		s += v
	}
	return s
}

func appendBytes() string {
	b := make([]byte, 0, 100)
	for _, v := range list {
		b = append(b, v...) // HL
	}
	return string(b)
}

func useBytesBuff() string {
	b := bytes.NewBuffer(make([]byte, 0, 100))
	for _, v := range list {
		b.WriteString(v) // HL
	}
	return b.String()
}
