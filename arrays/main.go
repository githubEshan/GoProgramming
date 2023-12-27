package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := [4]int{1, 2, 3,4}
	for i, v := range arr {
		fmt.Println("index", i, "value:", v)
	}

	fmt.Println(strings.Repeat("#", 20))

	for i:=0; i <(len(arr)); i++{
		fmt.Println("index", i, "value ", arr[i])
	}

	key := [...]int{
		1: 5,
		6:7,
	}
	fmt.Printf("%v\n", key)
	
}
