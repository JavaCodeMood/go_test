package main

import "fmt"

//切片可实现类似动态数组的功能

func main(){
	x := make([]int, 0, 5)   //创建容量为5的切片

	for i:=0;i<8;i++{  //给切片追加数据，当容量超出限制时，自动分配更大的存储空间
		x = append(x, i)
	}
	fmt.Println(x)

	//遍历切片
	for n := range x{
		fmt.Print(x[n],"-")
	}
}
