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

	c := map[string]string{
		"还有这种方法": "gpt害我",
	} //很神奇这边最后也要加逗号和json不一样
	fmt.Println(c)

	a = make(map[int]string) //重新make1遍。把a彻底清空。释放内存
	fmt.Println(a)

	//遍历for v,k:=range a,,二维字典什么的就不写了,,简单
}
