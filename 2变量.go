package main

import (
	"fmt"
)

var (
	全局变量  = "全局变量"
	全局变量2 = "太牛逼了"
)

func main() {
	var 测试 = "不指定类型"
	fmt.Println(测试)

	var age int = 18
	fmt.Println("age=", age)

	var age2 int
	fmt.Println("指定类型.不赋值,输出默认值", age2)

	age3 := "省略写法"
	fmt.Println(age3)

	fmt.Println("-----------------------------------------")
	var n1, n2, n3, n4 int
	fmt.Println(n1, n2, n3, n4)
	var n5, n6, n7, n8 = 999999, "jack", 2.21, "666"
	fmt.Println(n5, n6, n7, n8)
	n9, n0 := 3.3, "好好好"
	fmt.Println(n9, n0)
	fmt.Println(全局变量, 全局变量2)
	fmt.Println("-----------------------------------------")
	m1 := true
	fmt.Println(m1)
	fmt.Printf("查看变量类型%T", n5)
	字符串 := `
换行
用这个
反引号`

	字符串 += "然后" +
		"添加" + "要这样" +
		"+在一行末尾,不能是开头."
	fmt.Println(字符串)

	fmt.Println("-----------------------------------------")

	s1 := 65
	s2 := int8(s1) + 127
	fmt.Println(s2, "+128就会报错")
}
