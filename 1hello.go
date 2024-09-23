package main

//行注释
/*块注释*/
import "fmt"

func main() {
	fmt.Print("hello,,你好不换行")
	var age = 6 + 8
	fmt.Println(age)
	fmt.Println("1111", "22222", "3333",
		"4444", "5555", "逗号有空格")
	fmt.Print("1111", "22222", "3333",
		"4444", "5555", "逗号没有空格", "\n")
}
