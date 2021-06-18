package main

import "fmt"

func main() {
	var a int
	var f float32 = 12
	fmt.Println(a) // 0
	fmt.Println(f) // 12

	a = 5
	f = 14.0
	fmt.Println(a) // 5
	fmt.Println(f) // 14

	var i, j, k int = 1, 2, 3
	fmt.Println(i) // 1
	fmt.Println(j) // 2
	fmt.Println(k) // 3

	n := 100
	fmt.Println(n) // 100

	const cn int = 10
	const s string = "Hello"
	fmt.Println(cn) // 10
	fmt.Println(s)  // Hello

	const cn2 = 10
	const s2 = "Hello"
	fmt.Println(cn2) // 10
	fmt.Println(s2)  // Hello

	const (
		apple  = "Apple"
		banana = "Banana"
		orange = "Orange"
	)
	fmt.Println(apple)  // Apple
	fmt.Println(banana) // Banana
	fmt.Println(orange) // Orange

	const (
		zero = iota // 0
		one         // 1
		two         // 2
	)
	fmt.Println(zero)
	fmt.Println(one)
	fmt.Println(two)

}
