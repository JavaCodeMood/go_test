package main

import "fmt"

//使用cap函数返回切片的容量大小
//使用make函数创建一个切片

func main(){
	//使用make函数创建numbers切片数组
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	//缺省情况下声明没有输入切片，则将其初始化为nil，长度和容量均为0
	var numbers1 []int
	printSlice1(numbers1)
	if numbers1 == nil {
		fmt.Println("切片为空")
	}
}

//获取切片numbers的信息
func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}

func printSlice1(y []int){
	fmt.Printf("len %d cap = %d slice = %v\n", len(y), cap(y), y)
}
