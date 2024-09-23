package main

import "fmt"

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println(6)
		}
	}()
	str := "golang"
	fmt.Println(len(str))
	num1 := 1
	num2 := 0
	fmt.Println(num1 / num2)
	//不能直接1/0不然编译他不给你通过。

}
func main() {
	test()
	fmt.Println("okokok")
}
