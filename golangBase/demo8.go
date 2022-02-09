package main

import "fmt"

/**
  数组
*/
func main() {
	var a [2]int
	var b [2]int
	//数组长度不一样不可以直接比较
	// var b [1] int
	//a := [20] int {19:1} // 第20个赋值
	//a := [...]int{19:1}

	fmt.Println(a == b)
	fmt.Println(b)
}
