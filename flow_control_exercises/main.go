package main

import "fmt"

func main() {

	for i := 1; i <= 50; i++ {
		if i%7 == 0 { // if i is divisible by 7
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")
}


func flowControl(){
	count := 0
	for j:=0; j<=50;j++{
		if j%7 == 0 {
			continue
		}
		fmt.Printf("%d", j)
		count++
		if count ==3{
			break
		}
	}
	fmt.Println("")
}


func fix(){
	age := -9

	switch{
	case age < 0 || age > 100 :
		fmt.Println("Invalid Age")
	case age < 18 :
		fmt.Println("You are minor!")
	case age == 18 :
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}