package main

import "fmt"

func main() {
	//创建一个长度为1，类型为int的通道
	ch := make(chan int, 1)
	for i := 0; i < 4; i++ {
		select {
		case e, ok := <-ch:
			if !ok {
				fmt.Println("End.")
				return
			}
			fmt.Println(e)
			close(ch) //关闭通道
		default:
			fmt.Println("No Data!")
			ch <- 1
		}
	}
}
