package main

import "fmt"

//将字典类型内置，可直接从运行时层面获得性能优化

func main(){
	m := make(map[string]int)   //创建字典类型对象
	m["a"] = 100   //给字典添加键值对
	m["b"] = 200

	x, ok := m["b"]  //获取值
	fmt.Println(x,ok)

	delete(m, "a")    //根据键删除字典的值
	fmt.Println(m)
}
