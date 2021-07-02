package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum) // 55

	sum2 := 1
	for sum2 < 20 { // for문에서 조건문만 쓰면 while문처럼 동작
		sum2 *= 2
	}
	fmt.Println(sum2) // 32

	// for {
	// 	fmt.Println("무한 루프")
	// }

	number := []string{"one", "two", "three"}
	for idx, val := range number {
		fmt.Println(idx+1, val)
		// 1 one
		// 2 two
		// 3 three
	}

	num := 1
	for num < 20 {
		num++
		if num == 5 {
			num += 10
			fmt.Println(num) // 15
			continue         // for문으로 이동
		}
		if num > 18 {
			fmt.Println(num) // 19
			break            // for문 밖으로 이동
		}
	}
	if num == 20 {
		goto END
	}
END:
	fmt.Println("END") // END

	num2 := 0
BREAK1:
	for {
		if num2 == 0 {
			break BREAK1
		}
	}
	fmt.Println("Print!!!") // Print!!!
}
