package main

import "fmt"

func main() {
	var n1 = 7 + 9
	fmt.Println(n1)
	n2 := 99
	fmt.Println(n1 + n2)
	var s1 = "好" + "坏"
	fmt.Println(s1)

	fmt.Println(10 / 3)
	fmt.Println(10.0 / 3)
	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)
	n1++
	fmt.Println(n1)
}
