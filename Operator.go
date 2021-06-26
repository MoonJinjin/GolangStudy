package main

import "fmt"

func main() {
	var a int = 5 // 대입
	b := 2        // 선언 및 대입
	c := 1 + 2
	d := "Hello " + "World"

	fmt.Println(a) // 5
	fmt.Println(b) // 2
	fmt.Println(c) // 3
	fmt.Println(d) // Hello World

	// 산술연산자
	e := a - b
	f := a * b
	g := a / b
	h := a % b

	fmt.Println(e) // 3
	fmt.Println(f) // 10
	fmt.Println(g) // 2
	fmt.Println(h) // 1

	// 할당연산자
	a += 3
	fmt.Println(a) // 8
	d += "!!"
	fmt.Println(d) // Hello World!!
	a -= 1
	fmt.Println(a) // 7
	a *= 3
	fmt.Println(a) // 21
	a /= 2
	fmt.Println(a) // 10
	a %= 4
	fmt.Println(a) // 2

	i := -a
	j := +a

	fmt.Println(i) // -2
	fmt.Println(j) // 2

	// 관계연산자
	fmt.Println(3 == 3)     // true
	fmt.Println(5 == 2)     // false
	fmt.Println("A" == "A") // true
	fmt.Println("A" == "a") // false
	fmt.Println("A" != "A") // false

	fmt.Println(5 > 2)  // true
	fmt.Println(3 < 3)  // false
	fmt.Println(3 <= 3) // true

	// 논리연산자
	fmt.Println(!true)         //false
	fmt.Println(true && true)  // true
	fmt.Println(true && false) // false
	fmt.Println(true || true)  // true
	fmt.Println(true || false) // true

}
