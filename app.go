package main

import "fmt"


func main() {
multiplicity(5, 5)
}

func multiplicity(x, y int) {
	if x % y == 0 {
		fmt.Println(x, "кратен", y)
	} else {
		fmt.Println(x, "не кратен", y)
	}
}