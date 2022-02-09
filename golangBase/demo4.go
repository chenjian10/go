package main

import "fmt"

// 枚举和变量的顺序有关
const (
	a = 'A'
	b
	c = iota
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
