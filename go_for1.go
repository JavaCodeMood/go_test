package main

import "fmt"

func main() {
	map1 := map[int]string{1: "Golang", 2: "Java", 3: "go", 4: "python", 5: "Html5"}
	for key, value := range map1 {
		fmt.Printf("%d: %s\n", key, value)
	}

	fmt.Printf("\n\n")

	for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if v == 5 {
			continue
		}
		fmt.Printf("%d: %s\n", v, map1[v])
	}

	fmt.Printf("\n\n")
	for i := 10; i >= 0; i -= 2 {
		if i == 4 {
			break
		}
		fmt.Printf("%d\n", i)
	}

	fmt.Printf("\n\n")

	for j := 1; j <= len(map1); j++ {
		fmt.Printf("%d: %s\n", j, map1[j])
	}
}

