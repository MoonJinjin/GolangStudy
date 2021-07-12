package main

import "fmt"

func main() {
	PrintFunc() // Hello!!!!!

	a := "This is A"
	A(a) // This is A

	Variadic("A", "B", "C", "D", "E") // ABCDE

	Sum1 := Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(Sum1) // 55
}

func PrintFunc() {
	fmt.Println("Hello!!!!!")
}

func A(a string) {
	fmt.Println(a)
}

func Variadic(str ...string) {
	for _, s := range str {
		fmt.Print(s)
	}
	fmt.Println()
}

func Sum(nums ...int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
}
