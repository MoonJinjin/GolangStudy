package main

import "fmt"

func main() {
	str := `복수라인
	한줄
	두줄\n
	세줄`

	str2 := "단라인한줄두줄\n세줄"

	fmt.Println(str) // 복수라인
	//						한줄
	//						두줄\n
	//						세줄
	fmt.Println(str2) // 단라인한줄두줄
	// 세줄

	var i int = 10
	var u uint = uint(i)
	var f float32 = float32(i)

	fmt.Println(i) // 10
	fmt.Println(u) // 10
	fmt.Println(f) // 10

	str3 := "str"
	bytes := []byte(str3)
	str4 := string(bytes)

	fmt.Println(str3)  // str
	fmt.Println(bytes) // [115 116 114]
	fmt.Println(str4)  // str

}
