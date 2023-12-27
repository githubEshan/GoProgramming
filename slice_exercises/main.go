package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// 	Using a composite literal declare and initialize a slice of type string with 3 elements.
	// Iterate over the slice and print each element in the slice and its index.

	s := []string{"yo", "hello", "bye"}
	fmt.Println(s)

	//fixed the code
	mySlice := []float64{1.2, 5.6}

	mySlice[1] = 6

	a := 10
	mySlice[0] = float64(a)

	mySlice = append(mySlice, float64(a))

	fmt.Println(mySlice)

	// 	Declare a slice called nums with three float64 numbers.
	nums := []float64{2.3, 3.2, 5.9}
	fmt.Println(nums)
	// 2. Append the value 10.1 to the slice
	nums = append(nums, 10.1)
	fmt.Println(nums)
	// 3. In one statement append to the slice the values: 4.1,  5.5 and 6.6
	nums = append(nums, 4.1, 5.5, 6.6)
	fmt.Println(nums)
	// 4. Print out the slice
	// 5. Declare a slice called n with two float64 values
	n := []float64{2.2, 6.9}
	// 6. Append n to nums
	nums = append(nums, n...)
	// 7. Print out the nums slice
	fmt.Printf("The value of the slice nums is %v", nums)




// 	Create a Go program that reads some numbers from the command line and then calculates
// the sum and the product of all the numbers given at the command line.
// The user should give between 2 and 10 numbers.
	
	numbin := os.Args[1:]
	sum, product := 0., 1.


	for _, v := range numbin{
		num, err := strconv.ParseFloat(v, 64)
		_=err

		sum += num
		product *= num
	}
	fmt.Println(sum,product)
}
