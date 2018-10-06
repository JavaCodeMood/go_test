package main

import "fmt"

//可以直接调用匿名字段的方法，这种方式可以实现与继承类似的功能

type user struct {
	name string
	age int
}

func (u user) ToString() string{   //(u user)是ToString方法的接收者
	return fmt.Sprint("%+v",u)
}


type manager struct {
	user    //匿名嵌入其他类型
	title string
}

func main(){
	var m manager
	m.name = "断桥残雪"
	m.age = 18
	m.title = "闭月羞花"

	fmt.Println(m)
	fmt.Println(m.ToString())
}



