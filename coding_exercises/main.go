package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		a int
		b float64
		c bool
		d string
	)
	x, y, z := 20, 15.5, "Gopher!"

	_, _, _, _, _, _, _ = a, b, c, d, x, y, z

}

func test() {

	var a float64 = 7.1

	x, y := true, 3.7

	a, x = 5.5, false //fixed issue: can't have the short decleration for already declared variables

	_, _, _ = a, x, y
}

func test2() {

	version := "3.1"
	_ = version      //unused variable
	name := "Golang" //string requires double quotes, not single quotes
	fmt.Println(name)
}

//constants coding exercises

//Using the const keyword declare and initialize the following constants:
//1. daysWeek with value 7
//2. lightSpeed with value 299792458
//3. pi with value 3.14159

func constants() {
	const daysWeek int = 7
	const lightSpeed float64 = 299792458
	const pi float64 = 3.14159
}

func untypedConstants() {
	const daysWeek = 7
	const lightSpeed = 299792458
	const pi = 3.14159
}

//Calculate how many seconds are in a year.

func secondsInYear() {
	//1. Declare secPerDay constant and initialize it to the number of seconds in a day

	const (
		secPerDay = 3600 * 24
		daysYear  = 365
	)

	//3. Use fmt.Printf() to print out the total number of seconds in a year.
	fmt.Printf("Seconds in a year\n", secPerDay*daysYear)
}

func error() {

	const x int = 10
	var m = []int{1, 3, 4, 5, 6, 8}
	_ = m
}

func error2() {
	const a = 7
	const b float64 = 5.6
	const c = a * b

	x := 8
	_ = x

	//const noIPv6 = math.Pow(2, 128)
}

func print1() {

	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	fmt.Printf("%T%T%T", x, y, z)
	fmt.Printf("%v", z)
	fmt.Printf("%v", score)
}

func print4() {
	const x float64 = 1.422349587101
	fmt.Printf(".4%f", x)
}

func fixPrint() {

	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	circumference := 2 * pi * radius

	fmt.Printf("%q\n", shape)
	fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
	_ = shape
}

func conversions() {
	var (
		i int     = 3
		f float64 = 3.2
	)
	iconv := float64(i)
	fconv := int(f)
	_, _ = iconv, fconv

	//	Write a Go program that converts i to float64 and f to int.
}

func conversions1() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	_, _ = f, s1
	//i to string (int to string)
	s := strconv.Itoa(i)
	fmt.Printf("converted int to string %s", s)

	is, err := strconv.Atoi(s2)
	if err == nil {
		fmt.Printf("i type is %T, i value is %v\n", is, is)
	} else {
		fmt.Println("Can not convert string to int.")
	}
}

func alias1() {
	type duration int

	var hour duration

	fmt.Printf("%v", hour)
	fmt.Printf("%T", hour)

	hour = 3600
	fmt.Printf("%v", hour)
}

func aliasError() {

	type duration int
	var hour duration = 3600
	minute := 60
	fmt.Println(hour != duration(minute))
}
