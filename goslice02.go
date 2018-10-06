package main

import "fmt"

//切片允许指定上界和下界，以使用[lower-bound：upper-bound]获取它的子切片

func main(){
	//创建一个切片
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	fmt.Println("numbers == ", numbers)

	fmt.Println("numbers[1:4] == ", numbers[1:4])

	fmt.Println("numbers[:3] ==", numbers[:3])

	fmt.Println("numbers[4:] ==", numbers[4:])

	//子切片
	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	numbers2 := numbers[:5]
	printSlice(numbers2)

	numbers3 := numbers[2:6]
	printSlice(numbers3)

}

func printSlice(x []int){
	fmt.Printf("len = %d cap = %d slice = %v\n",len(x),cap(x),x)
}
