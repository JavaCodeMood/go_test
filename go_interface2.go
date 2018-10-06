package main

import "fmt"

type users struct {
	name string
	age int
}

func(u users) Print(){
	fmt.Println("%+v\n",u)
}

type Printer interface {   //接口类型
	Print()
}

func main(){
	var u users
	u.name = "婀娜多姿"  //只要包含接口所需要的全部方法，就表示实现了接口
	u.age = 22

	var p Printer = u
	p.Print()
}