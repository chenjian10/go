package main

import "fmt"

func main() {
	/*	f := closure(10)
		fmt.Println(f(1))
		fmt.Println(f(2))

		defer func() {

		}()*/
	A()
	B()
	C()
}

/**
  引用类型copy 地址 函数体里变 整体也变
  基本类型 copy 值
*/
func A() {
	fmt.Println("Func A")
}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")
}

func C() {
	fmt.Println("Func C")
}

// 闭包
func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
