package main

import "fmt"

func main() {
	a := 3
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three") // V
	}

	score := 80
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B") // V
	case score >= 70:
		fmt.Println("C") // Go는 자동으로 break를 해줌
	}

	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B") // V
		fallthrough      // break를 하지 않음
	case score >= 70:
		fmt.Println("C") // V
	}
}
