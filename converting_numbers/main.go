package main

import "fmt"

func main() {
	var x = 3
	var y = 3.1
	
	x =  int(float64(x)*y)
	fmt.Println(x)
}
