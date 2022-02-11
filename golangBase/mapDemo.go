package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[int]string)
	m[1] = "OK"
	delete(m, 1)
	fmt.Println(m)

	m2 := make(map[int]map[int]string)
	a, ok := m2[2][1] // ok 判断是否存在如果不存在创建
	fmt.Println(a, ok)
	if !ok {
		m2[2] = make(map[int]string)
	}
	m2[2][1] = "GOOD"
	a, ok = m2[2][1]
	fmt.Println(a, ok)

	sm := make([]map[int]string, 5)
	// 用 v 只是复制了map 不会影响原来的map
	for _, v := range sm {
		v = make(map[int]string, 1)
		v[1] = "OK"
		fmt.Println(v)
	}
	fmt.Println(sm)
	// 影响了原来的map
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
	// 实现key的排序
	m3 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	s := make([]int, len(m3))
	i := 0
	for k, _ := range m3 {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)
	// map[int]string ==>  map[string]int
	m4 := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(m4)
	m5 := make(map[string]int)
	for k, v := range m4 {
		m5[v] = k
	}
	fmt.Println(m5)

}
