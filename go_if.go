//go语言if条件语句
package main

import "fmt"

func main() {
	//定义局部变量并赋值
	var a int = 10

	/*使用if语句判断布尔表达式*/
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	}
	fmt.Printf("a的值为： %d\n", a)

}
