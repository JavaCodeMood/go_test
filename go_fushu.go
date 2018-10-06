//go语言复数
/*
go语言的复数类型有两个，即complex64和complex128。
存储这两个类型的值的空间分别需要8个字节和16个字节。
 复数类型的值一般由浮点数表示的实数部分、加号“+”、浮点数表示的虚数部分，以及小写字母“i”组成。
比如，3.7E+1 + 5.98E-2i
*/
package main

import "fmt"

func main() {
	var num1 = 3.7E+4 + 3.14E-2i
	var num2 = 6.18E-2 + 3.14E-2i

	fmt.Printf("浮点数 %E 表示的是：%f\n", num1, num1)
	fmt.Printf("浮点数 %E 表示的是：%f", num2, num2)
}
