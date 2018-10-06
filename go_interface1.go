package main

import "fmt"

//任何类型都实现了空接口
//接口只有方法声明，没有实现，没有数据字段
//接口可以匿名嵌入其他接口，或嵌入到结构中
//空接口可以作为任何类型数据的容器
//只有当接口存储的类型和对象都为nil时，接口才等于nil
//接口调用不会做receiver的字典转换
//接口是一个或多个方法签名的集合

//定义一个接口
type USB interface {
	Name() string
	Connect   //嵌入接口
}

type Connect interface {
	Connect()
}

//实现USB接口
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string{
	return pc.name
}

func (pc PhoneConnecter) Connect(){
	fmt.Println("connect:",pc.name)
}

func main(){
	var a USB
	a = PhoneConnecter{"Phone"}
	a.Connect()
	Disconnect(a)
	Disconnect1(a)
	fmt.Println("---------------")
	var a1 interface{}
	fmt.Println(a1 == nil)

	var p *int = nil
	a1 = p
	fmt.Println(a == nil)
}

func Disconnect(usb USB){
	if pc,ok := usb.(PhoneConnecter);ok {
		fmt.Println("Disconnected:",pc.name)
		return
	}
	fmt.Println("go接口实例")
}

func Disconnect1(usb interface{}){
	switch v:= usb.(type) {
	case PhoneConnecter:
		fmt.Println(v.name)
	default:
		fmt.Println("使用空接口")
	}
}
