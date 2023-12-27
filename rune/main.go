package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "ţară"

	fmt.Println("the decoded string value is ", str[3])

	for i:= 0; i<len(str); i++{
		fmt.Printf("%c", str[i])
	}
		fmt.Println("\n" + strings.Repeat("#", 20))
		
	for i:=0; i<len(str);{ r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i+=size
	}
}