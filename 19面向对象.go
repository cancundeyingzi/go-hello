package main

import "fmt"

type Teacher struct {
	Name string
	Age  int
	sex  string
}

func main() {
	var t1 Teacher
	t1.sex = "女"
	t1.Age = 18
	t1.Name = "蔡徐坤"
	fmt.Println(t1)

	var t2 Teacher = Teacher{"赵丽颖", 22, "女"}
	fmt.Println(t2)

	var t3 *Teacher = new(Teacher)
	(*t3).Age = 18
	(*t3).Name = "马牛逼"
	t3.sex = "男"
	fmt.Println(*t3)

	var t4 *Teacher = &Teacher{}
	(*t4).Age = 18
	(*t4).Name = "马牛逼"
	t4.sex = "男"
	fmt.Println(*t4)
}
