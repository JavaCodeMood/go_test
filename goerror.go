package main

import (
	"errors"
	"math"
	"fmt"
)

/*
Go编程提供了一个非常简单的错误处理框架，以及内置的错误接口类型，如下声明：

type error interface {
   Error() string
}
Go
函数通常返回错误作为最后一个返回值。 可使用errors.New来构造一个基本的错误消息
*/


func Sqrt(value float64) (float64, error){
	if value < 0 {
		return 0, errors.New("错误的参数")
	}
	return math.Sqrt(value), nil
}

func main(){
	result, err := Sqrt(-1)
	if err != nil {
		fmt.Println("发生错误1")
	}else{
		fmt.Println(result)
	}

	result, err = Sqrt(9)

	if err != nil {
		fmt.Println("发生错误2")
	}else{
		fmt.Println(result)
	}
}


