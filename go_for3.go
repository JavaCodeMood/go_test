package main

import (
	"fmt"
)

func main() {
	map1 := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}

	for key, value := range map1 {
		fmt.Printf("%d: %s\n", key, value)
	}

	fmt.Printf("\n")

	for _, v := range []int{1, 2, 3, 4} {
		fmt.Printf("%d: %s\n", v, map1[v])
	}

	fmt.Printf("\n")

	for i := 1; i <= len(map1); i++ {
		fmt.Printf("%d: %s\n", i, map1[i])
	}
}
