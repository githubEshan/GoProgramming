package main

import "fmt"

func main() {
	tp := [4]string{"Federer", "Nadal", "Djokovic", "Murray"}
	goat := [2]string{"Federer", "Sampras"}

inner:
	for index, name := range tp {
		for _, tp := range goat {
			if tp == name {
				fmt.Printf("found name %v at index %d", name, index)
				break inner
			}
		}
	}
}
