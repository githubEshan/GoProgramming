package main

import (
	"fmt"
)

func main() {
	// 1. Using the var keyword declare a string called name and initialize it with your name.
	var name string = "Eshan"
	// 2. Using short declaration syntax declare a string called country and assign the country you live in
	country := "US"
	// 3. Print the following string on multiple lines like this:
	// Your name: `here the value of name variable` Country: `here the value of country variable`
	fmt.Printf("Your name:%s\nhere the value of country variable%s\n", name, country)
	// a) He says: "Hello"
	fmt.Println("He says \"hello\"")
	// b) C:\Users\a.txt
	fmt.Println("C:\\Users\\a.txt")

// 	Declare a rune called r that stores the non-ascii letter ă
r:= "ă"
// 2. Print out the type of r
fmt.Printf("%T", r)
// 3. Declare 2 strings that contains the values ma and m
var s1, s2 string = "ma", "m"
// 4. Concatenate the strings and the rune in a new string (the new string will hold the value mamă ).
ss:= s1 +s2 +string(r)
fmt.Println(ss)



s3 := "țară means country in Romanian"
// Iterate over the string and print out byte by byte
for i:=0; i<len(s3); i++{
	fmt.Printf("%v", s3[i])
}
// 2. Iterate  over the string and print out rune by rune and the byte index where the rune starts 

//in the strings3 := "țară means country in Romanian"

}
