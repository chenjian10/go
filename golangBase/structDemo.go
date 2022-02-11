package main

import "fmt"
/**
   匿名结构必须顺序相同 
 */
type person struct {
	Name string
	Age int
	Contact struct{
		Phone, City string
	}
}

func main() {
	a := &person{
		Name: "lisi",
		Age: 19,
	}
	a.Contact.Phone = "12345678965"
	a.Contact.City = "济南"
	//a.Name = "zhangsan"
	//a.Age = 19
	AA(a)
	BB(a)
	fmt.Println(a)
	
	b := struct {
		Name string
		Age int
	}{
		Name: "joe",
		Age: 19 ,
	}
	
	fmt.Println(b)
	
	
}

func AA(per *person)  {
	per.Age = 13
	fmt.Println("A :" ,per)
}

func BB(per *person)  {
 	per.Age = 15
 	fmt.Println("B :" ,per)
}
