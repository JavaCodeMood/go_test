package main

import "fmt"

type MM int

func main(){
	var a MM
	a.Method()
	(*MM).Method(&a)

	var b MM
	b.Method1(100)
	fmt.Println(b)
}

func (a *MM) Method(){
	fmt.Println("Go语言方法")
}

func (b *MM) Method1(num int){
	*b += MM(num)   //类型转换
}