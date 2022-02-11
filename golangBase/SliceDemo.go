package main

import "fmt"

/**
其本身并不是数组，它指向底层的数组
作为变长数组的替代方案，可以关联底层数组的局部或全部
为引用类型
可以直接创建或从底层数组获取生成
使用len()获取元素个数，cap()获取容量
一般使用make()创建
如果多个slice指向相同底层数组，其中一个的值改变会影响全部

make([]T, len, cap)
其中cap可以省略，则和len的值相同
len表示存数的元素个数，cap表示容量

*/
func main() {
	var s1 []int
	fmt.Println(s1)

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	s2 := a[5:10]
	fmt.Println(s2)

	s3 := a[5:]
	fmt.Println(s3)

	s4 := make([]int, 3, 10)
	fmt.Println(s4, len(s4), cap(s4))

	b := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	//c:= []byte{'a','b','c','d','e','f','g','h','i','j','k'}
	sa := b[2:5]
	sb := b[3:5]
	fmt.Println(sa)
	fmt.Println(sb)
	sc := sa[1:5]
	fmt.Println(sc)

	s5 := make([]int, 3, 6)
	fmt.Println("%p\n", s5)
	s5 = append(s5, 1, 2, 3)
	fmt.Printf("%v %p\n", s5, s5)
	s5 = append(s5, 4, 5, 6)
	fmt.Printf("%v %p\n", s5, s5)

	s6 := a[2:5]
	s7 := a[1:3]
	fmt.Println(s6, s7)
	// append 后指向新的地址，改变就不影响其他指向该地址
	//s7 = append(s7,1,1,1,1,1,1)
	s7[0] = 9
	fmt.Println(s6, s7)

	s8 := []int{1, 2, 3, 4, 5, 6}
	//s9 := []int{7,8,9,1,1,1,1,1,1,1}
	s10 := s8
	s10[0] = 9
	fmt.Println(s10)
	fmt.Println(s8)

}
