package main

import "fmt"

/**
跳出多层循环：
   break
   continue
   goto
*/
func main() {
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println(" break 跳出无限循环")
	//  continue  必须是有限循环
LABEL2:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue LABEL2
		}
	}
	fmt.Println("continue 跳出循环")

	/*LABEL3:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			goto LABEL3
		}
	}
	fmt.Println("goto")*/

}
