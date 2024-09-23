package main

import "fmt"

func main() {
	goto lable1 //尽量别用会让可读性变差
	fmt.Println(1)
	fmt.Println(2)
lable1:
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	return //结束当前函数///////////
}
