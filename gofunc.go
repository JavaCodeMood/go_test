package main

import (
	"errors"
	"fmt"
)

func div(a , b int) (int, error){
	if b == 0{
		return 0,errors.New("除数不能为0")
	}
	return a / b,nil
}

//函数时第一类型，可以作为参数或返回值
func test(x int) func(){  //返回函数类型
	return func(){   //匿名函数
		println(x)   //闭包
	}
}

func test1(n, m int){
	defer fmt.Println("用defer定义延迟调用，无论函数是否出错，它都确保结束前被调用")

	println(n / m)
}

func main(){
	a,b := 10,2  //定义多个变量
	c,err := div(a,b)
	fmt.Println(c,err)

	x:= 100
	f := test(x)
	f()

	test1(10,0)
}
