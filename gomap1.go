package main

import "fmt"

//delete()函数用于从映射中删除项目。它需要映射以及指定要删除的相应键。

func main (){
	//创建一个map
	colorMap := map[string] string {"红色":"red","黄色":"yellow","绿色":"green","黑色":"black","蓝色":"blue"}
	fmt.Println("map==", colorMap)

	//遍历map
	for country := range colorMap{
		fmt.Println(country, "->", colorMap[country])
	}

	fmt.Println("\n")

	//根据键删除某个值
	delete(colorMap, "黑色")
	for count := range colorMap {
		fmt.Println(count, "->", colorMap[count])
	}

}
