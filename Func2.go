package main

import "fmt"

func main() {
	// 익명함수
	func() {
		fmt.Println("abc")
	}() // abc

	func(a int) {
		fmt.Println(a * a)
	}(10) // 100

	// Func1 변수에 익명함수 대입
	Func1 := func(a int, b int) int {
		return a + b
	}
	fmt.Println(Func1(5, 10)) // 15

	// GoLang의 함수는 일급함수이다.
	// 함수가 다른 함수의 파라미터로 전달 가능
	add := Any(Func1, 100, 200)
	fmt.Println(add) // 300

	minus := Any(func(minus1 int, minus2 int) int {
		return minus1 - minus2
	}, 100, 50)
	fmt.Println(minus) // 50

}

func Any(f func(int, int) int, a int, b int) int {
	return f(a, b)
}

// 함수 원형 정의
type FuncType func(int, int) int

func Any2(f FuncType, a int, b int) int {
	return f(a, b)
}
