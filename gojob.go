package main

import (
	"fmt"
	"time"
)

//go并发

func task(id int){
	for i:=0;i<10;i++{
		fmt.Printf("%d : %d\n",id, i)
		time.Sleep(time.Second)
	}
}

func main(){
	go task(1)   //创建goroutinne
	go task(2)
	go task(3)

	time.Sleep(time.Second * 6)
}
