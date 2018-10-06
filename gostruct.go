package main

import "fmt"

//go语言结构体可以匿名嵌入其他类型

type user struct {
	name string
	age int
}

type manager struct {
	user     //匿名嵌入其他类型
	title string
}

func main(){
	var m manager

	m.name = "liuhefei"
	m.age = 22
	m.title = "没了你，万杯觥筹不过是提醒寂寞"
	fmt.Println(m)
}
