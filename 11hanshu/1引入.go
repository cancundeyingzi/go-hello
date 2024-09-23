package main

import (
	"fmt"
)

func cal(num1 int, num2 int) (int, int) { //最后的int是返回值类型(sum/r的类型
	sum := 0
	sum = num1 + num2
	r := num1 - num2
	return sum, r
}

func cal2(num1 int, num2 int) (sum int, r int) { //返回值也能这么写
	sum = 0
	sum = num1 + num2
	r = num1 - num2
	return
}
func test(args ...int) { //无限int参数
	fmt.Println(args)
	fmt.Println(args[2])
}

func test2(num *int) { //使用指针,当全局变量用
	*num = 999999
}

func test3(函数 func(num *int)) { //函数也能当参数传进来
	sum := 1 //这样子改不了外面的sum,,因为没传进来
	函数(&sum)
}

func main() {
	sum, r := cal(10, 20)
	fmt.Println(sum, r)
	test(666, 5445, 464, 654, 6546, 464, 164, 64, 16, 416, 416, 41)
	//test2(&sum)
	fmt.Println(sum, &sum)

	a := test
	a(10, 20, 50, 60) //等于test(10,20,50,60)

	fmt.Printf("%T\n", cal) //cal变量的类型
	test3(test2)            //不能在这里给test2参数,要写在29行test3里面
	fmt.Println("最后", sum, &sum)
}
