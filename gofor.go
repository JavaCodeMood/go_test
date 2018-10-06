package main

import "fmt"

//go流程控制语句

func main(){
	x := []int{100,200,300,400,500}
	for i,n := range x {
		println(i,",",n)
	}

	a := 100
	switch  {
	case a > 0:
		fmt.Println("a = ",a)
	case a < 0:
		fmt.Println("a = ",-a)
	default:
		print("0")

	}
}
