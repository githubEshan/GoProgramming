package main

import "fmt"

// 1. Using the var keyword, declare an array called cities with 2 elements of type string. Don't initialize the array.
// Print out the cities array and notice the value of its elements.
// 2. Declare an array called grades of type [3]float64 and initialize only the first 2 elements using an array literal.
// Print out the grades array and notice the value of its elements.
// 3. Declare an array called taskDone using the ellipsis operator. The elements are of type bool. Print out taskDone.
// 4. Iterate over the array called cities using the classical for loop syntax and the len function and print out element by element.
// 5. Iterate over grades using the range keyword and print out element by element.
func main(){
	var cities = [2]string{}
		fmt.Println(cities)


	grades := [3]float64{
		0: 2.3,
		1:4.3,
	}
		fmt.Println(grades)

		for i,v := range grades{
			fmt.Println(i,v)
		}

	taskDone := [...]bool{}
	fmt.Println(taskDone)

	for j:=0; j< len(cities); j++{
		fmt.Println(j)
	}
	exercise2()
}

func exercise2(){
	// Write a Go program that counts the number of positive even numbers in the array.
	nums := [...]int{30, -1, -6, 90, -6}
	count := 0

	for i:=0; i<len(nums);i++{
		if i%2 == 0{
			count++
		}
	}
	fmt.Printf("the number of even elenments in the array are %d", count)
}


func fixCode(){
	myArray := [3]float64{1.2, 5.6}
    myArray[2] = 6
    a := 10
    myArray[0] = float64(a)
    myArray[2] = 10.10
    fmt.Println(myArray)
}