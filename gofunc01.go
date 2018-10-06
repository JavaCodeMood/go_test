package main

import "fmt"

/*
Go编程语言支持递归，即函数调用自身的函数。 但是在使用递归时，程序员需要注意在函数中定义或设置一个退出条件，否则它会进入无限循环。

递归函数非常有用，可用于解决许多数学问题，如计算数字的阶乘，生成斐波那契数列等。
*/

func factorial(i int) int {
	if i<= 1 {
		return 1
	}
	return i * factorial(i -1)
}

func main(){
	var i int = 15
	fmt.Printf("斐波那契数列第 %d 个数是 %d", i, factorial(i))

	fmt.Println("\n")
	for x := 0;x <= 15;x++ {
		fmt.Println("第", x ,"个斐波那契数列是：",factorial(x))
	}
}

