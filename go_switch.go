package main

import (
	"fmt"
	"math/rand"
)


func main() {
	ia := []interface{}{byte(6), 'a', uint(10), int32(-4)}
	switch v := ia[rand.Intn(4):4]; interface{}(v).(type) {
	case []interface{}:
		fmt.Println("类型switch语句")
	case byte:
		fmt.Println("数据类型为byte")
	default:
		fmt.Println("如果前面条件不满足将会执行这条")
	}

	age := 40
	switch age {
	case 10:
		fmt.Println("小孩")
	case 20:
		fmt.Println("少年")
	case 30:
		fmt.Println("青年")
	case 40:
		fmt.Println("中年")
		fallthrough
	case 50:
		fmt.Println("中年后")
	case 60:
		fmt.Println("古稀")
	default:
		fmt.Println("不是人类")
	}
}

