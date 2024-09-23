package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("hello")
	fmt.Println("看执行顺序")
	fmt.Println("'栈,倒序,,函数执行完后再执行defer'")
	//但是它的变量不会跟着变化。比如最开始设置a等于零。然后他这个函数输出a..........后面a等于a加一。此时运行会显示零而不是一.
}
