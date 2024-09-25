package main

import "fmt"

// 只有字典创建的时候需要用make函数初始化。
func main() {
	var a map[int]string
	a = make(map[int]string)
	a[1213] = "啊啊啊"
	fmt.Println(a)

	b := make(map[string]string)
	b["好好好"] = "对对对"
	fmt.Println(b)
}
