package main

import (
	"math"
	"fmt"
)

//Go编程提供了另一种称为接口(interfaces)的数据类型，它代表一组方法签名。
// struct数据类型实现这些接口以具有接口的方法签名的方法定义。

//定义一个接口
type Shape interface {
	area() float64
}

//定义一个圆
type Circle struct {
	x,y,radius float64
}

//定义一个矩形
type Rectangle struct{
	width, height float64
}

//计算圆的面试、
func (circle Circle)area() float64{
	return math.Pi * circle.radius * circle.radius
}

//计算矩形的面积
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

//定义一个形状，用于获取面试
func getArea(shape Shape) float64 {
	return shape.area()
}

func main(){
	//圆的参数
	circle := Circle{x:0,y:0,radius:6}
	//矩形的参数
	rectangle := Rectangle{width: 5, height: 7}
	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
	fmt.Printf("圆周率：%f\n", math.Pi)
}
