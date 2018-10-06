package main

import "fmt"

func main(){
	//创建一个切片
	numbers := []int{0,1,2,3,4,5,6,7,8,9}
	//遍历这个切片
	for i := range numbers {
		fmt.Println("第" ,i, "个切片元素：", numbers[i])
	}

	//创建一个map
	countryCapitalMap := map[string] string{"红色":"red","黑色":"black","黄色":"yellow","绿色":"green","蓝色":"blue","紫色":"pink"}
	//遍历map
	for country := range countryCapitalMap {
		fmt.Println("数组元素：",country,"是", countryCapitalMap[country])
	}

	//打印map的键值对
	for key,value := range countryCapitalMap {
		fmt.Println(key, "->", value)
	}
}
