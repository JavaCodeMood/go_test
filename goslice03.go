package main

import "fmt"

//切片(Slice)允许使用append()函数增加切片的容量(大小)。
// 使用copy()函数，将源切片的内容复制到目标切片

func main(){
	var numbers []int
	printSlice(numbers)
	fmt.Println("numbers1 ==" , numbers)

	//向切片中追加元素
	numbers = append(numbers,1)
	numbers = append(numbers,4)
	numbers = append(numbers,8)
	printSlice(numbers)
	fmt.Println("numbers2 ==", numbers)

	numbers = append(numbers, 100,200,300)
	printSlice(numbers)
	fmt.Println("numbers3 ==", numbers)

	numberss := make([]int, len(numbers), (cap(numbers)) * 3)
	//将numbers切片中的数据复制到numberss中
	copy(numberss, numbers)
	printSlice(numberss)

}

func printSlice(x []int){
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}