package main

import "fmt"

type Students struct {
	Name string
	Age  int
}

func (s *Students) String() string {
	str := fmt.Sprintln("自动调用String()方法", s.Age, s.Name)
	return str
}

func main() {
	stu := Students{
		Name: "丽丽",
		Age:  9,
	}
	fmt.Println(&stu) // 记得传内存地址。这样就会自动调用上面的string方法。
}
