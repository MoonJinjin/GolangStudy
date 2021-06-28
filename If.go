package main

import "fmt"

func main() {
	num := 1

	if num == 1 {
		fmt.Println("one")
	} // one

	num = 2

	if num == 1 {
		fmt.Println("one")
	} else if num == 2 {
		fmt.Println("two")
	} // two

	num = 3

	if num == 1 {
		fmt.Println("one")
	} else if num == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("other")
	} // other

	if num = 5; num == 5 {
		fmt.Println("five")
	} // five
}
