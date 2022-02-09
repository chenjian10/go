package main

import "fmt"

/*
  6: 0110
  11: 1011
 ----------
  &  0010 = 2    （相同）
  |  1111 = 15  （结合）
  ^  1101 = 13  (不同为1)
  &^ 0100 = 4   （后面为 1 强制转化为 0）

==================
<<  位移运算符 左移
>>   位移运算符  右移
*/
func main() {
	fmt.Println(1 << 10)
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)

}
