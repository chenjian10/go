package main

import "fmt"

func main() {
	// 指针
	a := 10
	var p *int = &a
	fmt.Println(*p)

	// if 判断
	if a := 1; a > 0 {
		fmt.Println(a)
	}
	fmt.Println(a)

	//for 语句
	// 第一种 where(true)
	for {
		a++
		if a > 13 {
			break
		}
		fmt.Println(a)
	}
	// 第二种
	/*	for a < 3 {

		}*/
	//第三种
	/*	for i := 0; i < 3; i++ {

		}*/

	// swich 语句
	switch a := 1; {
	case a >= 0:
		fmt.Println("a >= 0")
		fallthrough
	case a >= 1:
		fmt.Println("a >= 1")
	default:
		fmt.Println("None")
	}
	fmt.Println(a)

}
