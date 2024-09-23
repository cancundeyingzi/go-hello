package main

import "fmt"

func main() {
	//方式1
	var age int
	fmt.Println("输入年龄")
	fmt.Scanln(&age) //必须用指针,,是ln
	fmt.Println(age)

	var name string
	var score float32
	fmt.Println("录入年龄、姓名、成绩，使用空格进行分隔")
	fmt.Scanf("%d %s %f", &age, &name, &score) //int,字符串,浮点,是f
	fmt.Println(age, name, score)
}
