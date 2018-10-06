package main

//映射(Map)，它将唯一键映射到值。 键是用于在检索值的对象。
// 给定一个键和一个值就可以在Map对象中设置值。设置存储值后，就可以使用其键检索它对应的值了。

import "fmt"

func main(){
	//定义一个map变量
	var countryCapitalMap map[string]string
	//创建map
	countryCapitalMap = make(map[string]string)
	//设置键值对
	countryCapitalMap["红色"] = "red"
	countryCapitalMap["黑色"] = "balck"
	countryCapitalMap["黄色"] = "yellow"
	countryCapitalMap["绿色"] = "green"
	countryCapitalMap["蓝色"] = "blue"

	//遍历map集合
	for country := range countryCapitalMap {
		fmt.Println("键：", country , "值：", countryCapitalMap[country])
	}

	//判读map中是否存在某个key-value对
	capital,ok := countryCapitalMap["红色"]
	if ok {
		fmt.Println("key存在map中,值为", capital)
	}else{
		fmt.Println("key不存在map中")
	}
}
