package main

import (
	"fmt"
	"reflect"
)

type User struct {
     Id int
     Name string
     Age int
}
type Manager struct {
	User
	title string
}
 // 利用反射调用方法
func (u User) Hello(name string)  {
	fmt.Println("Hello ", name,", my name is", u.Name)
}

func main() {
    u := User{1, "OK", 12}
    info(u)

    m := Manager{User:User{1,"ok",12}, title: "12"}
    t := reflect.TypeOf(m)
    fmt.Printf("%#v\n", t.FieldByIndex([]int{0,0}))
    
    x := 123
    v := reflect.ValueOf(&x)
    v.Elem().SetInt(999)
    fmt.Println(x)

	Set(&u)
	fmt.Println(u)

	v1 := reflect.ValueOf(u)
	mv := v1.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)

}

func info(o interface{})  {
	t := reflect.TypeOf(o)
	fmt.Println("type: ", t.Name())
	// 错误的时候
	if k := t.Kind(); k != reflect.Struct{
		fmt.Println("XXXX")
		return
	}
	v := reflect.ValueOf(o)
	fmt.Println("Fields:",v)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:  %v= %v\n",f.Name,f.Type,val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n",m.Name,m.Type)
	}
}
// 改变某个属性值 反射
func Set(o interface{})  {
    v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXXX")
		return
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
		
	}
	if  f.Kind() == reflect.String {
		f.SetString("BAYBAY")
	}
	
	
	
}
