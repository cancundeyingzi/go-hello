package main

import "fmt"

type Student struct {
	Age int
}

type Person struct {
	Age int
}

func main() {
	s := Student{10}
	p := Person{15}
	s = Student(p) //想要这么操作得保证person和student里面的变量一样。比如说这里都只有一个age int...东西
	fmt.Println(s)
}
