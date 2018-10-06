package main

import "fmt"

//类型转换
//类型转换是一种将变量从一种数据类型转换为另一种数据类型的方法。
//可以使用转换操作符将值从一种类型转换为另一种类型，如下所示： type_name(expression)

func main(){
	var a int = 128
	var b float32 = 3.14
	var c string = "3.456"

	d := (float32(a)) + b
	fmt.Println("类型转换后：d = ", d)

	e := (string(c))
	fmt.Println("float类型转化为字符串类型：", e)

}
