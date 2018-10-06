//Go 语言的 goto 语句可以无条件地转移到过程中指定的行。
package main

import "fmt"

func main() {
	//定义局部变量
	var a int = 10

	//循环
Loop:
	for a < 20 {
		if a == 15 {
			//跳过迭代
			a = a + 1
			goto Loop
		}
		fmt.Printf("a的值为：%d\n", a)
		a++
	}
}
