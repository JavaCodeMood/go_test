package main

import "fmt"

// 实现斐波那契数列

func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i - 1) + fibonaci(i - 2)
}

func main(){
	var i int
	for i = 0; i<= 10;i++{
		fmt.Printf("%d ", fibonaci(i))
	}

}
