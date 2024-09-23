package main

import "fmt"

func main() {

	a := []int{1, 2, 3}
	var b [2]int
	c := []int{2: 66, 0: 100}
	b[0] = 1
	b[1] = 2
	d := [6][8]int{}
	fmt.Println(a, b, c, d)

	for key, value := range a {
		fmt.Println(key, value)
	}
}
