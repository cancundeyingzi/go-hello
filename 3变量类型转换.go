package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 1
	var srca = fmt.Sprint(a)
	fmt.Println(srca)
	fmt.Printf("查看变量类型%T", srca)
	fmt.Println()

	var b = strconv.Itoa(a)
	fmt.Println(b)
	fmt.Printf("查看变量类型%T", b)
	fmt.Println()

	fmt.Println("---------------------------------")

	var n1 int = 18
	var s1 string = strconv.FormatInt(int64(n1), 10) //参数:第一个参数必须转为int64类型 ，第二个参数指定字面值的进制形式为十进制
	fmt.Printf("s1对应的类型是: %T ，s1 = %q \n", s1, s1)

	var n2 float64 = 4.29
	var s2 string = strconv.FormatFloat(n2, 'f', 9, 64)
	//第二个参数:'f’(-ddd.dddd) 第三个参数: 9 保留小数点后面9位 第四个参数:表示这个小数是float64类型
	fmt.Printf("s2对应的类型是: %T ，s2 = %q \n", s2, s2)

	var n3 bool = true
	var s3 string = strconv.FormatBool(n3)
	fmt.Printf("s3对应的类型是: %T ，s3 = %q \n", s3, s3)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------")

	var ss1 string = "true" //除了true和1,别的应该都是默认值(false)
	var bb bool
	//ParseBool这个函数的返回值有两个:(value bool，err error) value就是我们得到的布尔类型的数据，err出现的错误 我们只关注得到的布尔类型的数据，err可以用 直接忽略
	bb, _ = strconv.ParseBool(ss1)
	fmt.Printf("b的类型是: %T,bb=%v \n", bb, bb)

	var ss2 string = "19" //不是数字一律默认值0
	var num1 int64
	num1, _ = strconv.ParseInt(ss2, 10, 64) //ss2转10进制,int64
	fmt.Printf("num1的类型是: %T,num1=%v \n", num1, num1)

	var ss3 string = "3.14" //不是数字一律默认值0
	var f1 float64
	f1, _ = strconv.ParseFloat(ss3, 64) //float64
	fmt.Printf("f1的类型是: %T,f1=%v \n", f1, f1)
}
