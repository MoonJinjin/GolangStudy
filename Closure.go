package main

import "fmt"

func nextValue() func() int {
	i := 0
	return func() int {
		i++ // 함수의 바깥에 있는 변수 사용
		return i
	}
}
func main() {
	next := nextValue()
	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3

	// 새로운 클로저 함수를 생성하면 i는 0으로 초기화된다.
	next2 := nextValue()
	fmt.Println(next2()) // 1
	fmt.Println(next2()) // 2
	fmt.Println(next2()) // 3

	// 기존 클로저 함수의 i는 그대로 유지
	fmt.Println(next()) // 4
	fmt.Println(next()) // 5
}
