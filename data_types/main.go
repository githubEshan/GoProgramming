package main

import "fmt"

func main(){

	var i uint32 = 0 //for zero and positive n
	var n int32 = 32
	
	fmt.Printf("%T\n%T\n", i, n)
	
	//arrays 
	var numb = [4]int{1,2,3,4}
	fmt.Print(numb)

	//map
	balances := map[string]float32{
	}
	fmt.Printf("%T\n", balances)


	//structure

	type Person struct{
		name string 
		age int
	}

	var you Person
	fmt.Printf("%T\n", you)

	//pointers
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the vaue of %v\n", ptr, ptr)
}