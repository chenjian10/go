package main

import "fmt"

type (
	byte     int8
	rune     int32
	ByteSize int64
)

func main() {
	var a, b, c, d = 1, 2, 3, 4
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
