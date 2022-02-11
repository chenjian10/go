package main

import "fmt"

type human struct {
	Sex int
}
type teacher struct {
	human
	Name string
	Age int
}

type student struct {
	human
	Name string
	Age int
}

func main() {
	a:= teacher{Name: "joe",Age: 19,human: human{Sex: 0}}
	b:= student{Name: "joe",Age: 19,human: human{Sex: 0}}
	a.Age = 25
	a.human.Sex = 100
	fmt.Println(a,b)
}
